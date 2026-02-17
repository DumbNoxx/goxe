package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
	rTime "github.com/DumbNoxx/goxe/internal/processor/reportTime"
	"github.com/DumbNoxx/goxe/internal/utils"
)

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
