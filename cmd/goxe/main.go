package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/ingestor"
	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/internal/processor"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
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

func viewConfig() {
	dir, _ := os.UserConfigDir()
	configPath := filepath.Join(dir, "goxe", "config.json")
	initialStat, err := os.Stat(configPath)
	if err != nil {
		log.Fatal(err)
	}
	lastModified := initialStat.ModTime()

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		currentStat, err := os.Stat(configPath)
		if err != nil {
			log.Fatal(err)
		}
		if currentStat.ModTime().After(lastModified) {
			fmt.Println("Config update, reload...")
			lastModified = currentStat.ModTime()
			options.Config = options.ConfigFile()
		}
	}
}

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	var wg sync.WaitGroup
	pipe := make(chan *pipelines.LogEntry, 100)
	var mu sync.Mutex

	options.CacheDirGenerate()

	wg.Add(1)
	go processor.Clean(pipe, &wg, &mu)
	go ingestor.Udp(pipe, &wg)
	go viewConfig()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(pipe)

	wg.Wait()

}
