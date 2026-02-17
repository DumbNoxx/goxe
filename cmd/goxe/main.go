package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"slices"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/DumbNoxx/goxe/internal/ingestor"
	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
	rTime "github.com/DumbNoxx/goxe/internal/processor/reportTime"
	"github.com/DumbNoxx/goxe/internal/utils"
	pkg "github.com/DumbNoxx/goxe/pkg/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	versionFlag *bool
	isUpgrade   *bool
	version     string
)

func init() {
	versionFlag = flag.Bool("v", false, "")
	isUpgrade = flag.Bool("is-upgrade", false, "Internal use for hot-swap")
}

func executeHandoff(once *sync.Once, cancel context.CancelFunc, pipe chan<- *pipelines.LogEntry, wgProc, wgProd *sync.WaitGroup) {
	once.Do(func() {
		cancel()

		done := make(chan struct{})
		go func() {
			wgProd.Wait()
			close(done)
		}()

		select {
		case <-done:
		case <-time.After(2 * time.Second):
			fmt.Println("[System] Handoff: Force closing producers...")
		}

		close(pipe)
		wgProc.Wait()
	})
}

func getVersion() string {
	if version != "" {
		return version
	}

	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			return info.Main.Version
		}
	}
	return "vDev-build"
}

func viewConfig(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	dir, _ := os.UserConfigDir()
	configPath := filepath.Join(dir, "goxe", "config.json")
	initialStat, err := os.Stat(configPath)
	if err != nil {
		log.Fatal(err)
	}
	lastModified := initialStat.ModTime()

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			currentStat, err := os.Stat(configPath)
			if err != nil {
				log.Fatal(err)
			}
			if currentStat.ModTime().After(lastModified) {
				fmt.Println("Config update, reload...")
				lastModified = currentStat.ModTime()
				options.Config = options.ConfigFile()
				utils.TimeReportFile = utils.UserConfigHour()
				rTime.GetReportFileTime()
				rTime.GetReportPartialTime()
				filters.LoadFiltersWord()
			}
		case <-ctx.Done():
			return
		}
	}
}

func getVersionLatest(req *http.Request, res *http.Response, ctx context.Context) (response pkg.ResponseGithubApi) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com/repos/DumbNoxx/goxe/releases/latest", nil)
	if err != nil {
		log.Println(err)
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to unmarshal github response:", err)
	}
	return response
}

func viewNewVersion(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		res            *http.Response
		req            *http.Request
		currentVersion = getVersion()
	)

	ticker := time.NewTicker(60 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			release := getVersionLatest(req, res, ctx)

			if release.Tag_name == "vDev-build" {
				continue
			}

			if release.Tag_name == currentVersion {
				continue
			}

			fmt.Printf("Update available: %s -> %s\n", currentVersion, release.Tag_name)

			fmt.Println("--- Release Notes ---")
			fmt.Printf("\n%v\n", release.Body)
			fmt.Println("----------------------")
		case <-ctx.Done():
			return
		}
	}
}

func autoUpdate(ctx context.Context, cancel context.CancelFunc, pipe chan<- *pipelines.LogEntry, wgProcessor, wgProducer *sync.WaitGroup, once *sync.Once) {
	var (
		req *http.Request
		res *http.Response
	)
	currentLocation, err := os.Executable()
	home, _ := os.UserHomeDir()
	gopath := filepath.Join(home, "go")
	version := getVersionLatest(req, res, ctx)
	v := getVersion()
	if err != nil {
		log.Fatal(err)
	}
	if version.Tag_name != v {
		if !strings.HasPrefix(currentLocation, gopath) {
			fmt.Println("[Test] Local build detected, recompiling...")
			tempBin := currentLocation + ".tmp"
			cmd := exec.Command("go", "build", "-o", tempBin, "./cmd/goxe")
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Build failed: %v\n", err)
				log.Printf("Compiler says: %s\n", string(output))
				return
			}
			if err := os.Rename(tempBin, currentLocation); err != nil {
				fmt.Printf("[Error] Failed to swap binary: %v\n", err)
				return
			}

			fmt.Println("[System] Preparing handoff, flushing buffers...")
			executeHandoff(once, cancel, pipe, wgProcessor, wgProducer)
			err = syscall.Exec(currentLocation, []string{currentLocation, "-is-upgrade"}, os.Environ())

			fmt.Printf("\n[Error] ¡El salto a V2 falló!: %v\n", err)
			os.Exit(1)

			<-ctx.Done()
			return
		}
		if strings.HasPrefix(currentLocation, gopath) {
			cmd := exec.Command("go", "install", "github.com/DumbNoxx/goxe/cmd/goxe@latest")
			err := cmd.Run()
			if err != nil {
				log.Println(err)
				return
			}
		}
		fmt.Println("[System] Preparing handoff, flushing buffers...")

		executeHandoff(once, cancel, pipe, wgProcessor, wgProducer)
		err = syscall.Exec(currentLocation, []string{currentLocation, "-is-upgrade"}, os.Environ())

		fmt.Printf("\n[Error] ¡El salto a V2 falló!: %v\n", err)
		os.Exit(1)

		<-ctx.Done()
		return
	}

	if strings.HasPrefix(currentLocation, "/usr/bin/goxe") {
		fmt.Println("Goxe was installed via a package manager. Please use your package manager to update it to avoid versioning conflicts.")
	}
	<-ctx.Done()
}

func main() {
	flag.Parse()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	stopChan := make(chan os.Signal, 1)
	pipe := make(chan *pipelines.LogEntry, 100)
	var (
		wgProcessor sync.WaitGroup
		wgProducer  sync.WaitGroup
		mu          sync.Mutex
		once        sync.Once
	)

	if *versionFlag {
		fmt.Println(getVersion())
		os.Exit(0)
	}
	if *isUpgrade {
		fmt.Println("[System] System updated")
	}
	arg := os.Args

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, watchSignals...)

	go func() {
		for sig := range sigChan {
			if isUpdateSignal(sig) {

				fmt.Println("\n[System] Update signal received! Starting auto-update...")
				ticker := time.NewTicker(1 * time.Second)
				count := 1
				updateDone := false
				defer ticker.Stop()
				for !updateDone {
					select {
					case <-ticker.C:
						if count != 5 {
							fmt.Printf("%d..", count)
						}
						if count == 5 {
							fmt.Printf("%d\n", count)
							fmt.Println("Updating...")
							autoUpdate(ctx, cancel, pipe, &wgProcessor, &wgProducer, &once)
							updateDone = true
						}
						count++
					case <-ctx.Done():
						return
					}

				}
				return
			}
			if sig == os.Interrupt {
				cancel()
			}
		}
	}()

	if slices.Contains(arg, "update") {
		fmt.Println("Sending update signal to the active instance...")
		cmd := exec.Command("pkill", "-SIGUSR1", "goxe")
		cmd.Run()
		return
	}

	options.CacheDirGenerate()

	wgProcessor.Add(1)
	go processor.Clean(ctx, pipe, &wgProcessor, &mu)
	wgProducer.Add(1)
	go ingestor.Udp(ctx, pipe, &wgProducer)
	wgProducer.Add(1)
	go viewConfig(ctx, &wgProducer)
	wgProducer.Add(1)
	go viewNewVersion(ctx, &wgProducer)

	signal.Notify(stopChan, os.Interrupt)
	<-stopChan

	fmt.Println("[System] Shutdown app, flushing buffers...")
	executeHandoff(&once, cancel, pipe, &wgProcessor, &wgProducer)

}
