package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/internal/ingestor"
	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	versionFlag *bool
	version     string
)

func init() {
	versionFlag = flag.Bool("v", false, "")
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
				filters.LoadFiltersWord()
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	var wgProcessor sync.WaitGroup
	var wgProducer sync.WaitGroup
	pipe := make(chan *pipelines.LogEntry, 100)
	var mu sync.Mutex
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	options.CacheDirGenerate()

	wgProcessor.Add(1)
	go processor.Clean(ctx, pipe, &wgProcessor, &mu)
	wgProducer.Add(1)
	go ingestor.Udp(ctx, pipe, &wgProducer)
	wgProducer.Add(1)
	go viewConfig(ctx, &wgProducer)

	<-ctx.Done()

	done := make(chan struct{})
	go func() {
		wgProducer.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		log.Println("[System] Force closing producers...")
	}

	close(pipe)
	wgProcessor.Wait()

}
