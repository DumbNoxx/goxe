package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"syscall"
	"time"

	pkg "github.com/DumbNoxx/goxe/pkg/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

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
