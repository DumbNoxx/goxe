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

// viewConfig monitors the configuration file for changes and reloads it automatically.
//
// Parameters:
//
//   - ctx: context for cancellation; when cancelled, the function terminates.
//   - wg: WaitGroup to notify the caller that the goroutine has finished (calls wg.Done() at the start).
//
// Returns:
//
//   - void: the function runs in an infinite loop until ctx is cancelled.
//
// The function performs:
//
//   - Retrieves the configuration file path: <UserConfigDir>/goxe/config.json.
//
//   - Reads the file's initial modification date using os.Stat.
//
//   - Creates a ticker that fires every second (polling).
//
//   - On every ticker tick:
//
//     -Re-checks the file's modification timestamp.
//
//     -If the timestamp is more recent than the last recorded one, the file has changed.
//
//     -Prints "Config update, reload..." to the console.
//
//     -Updates the global options.Config variable by calling options.ConfigFile().
//
//     -Recalculates utils.TimeReportFile using utils.UserConfigHour().
//
//     -Restarts the report tickers by calling rTime.GetReportFileTime() and rTime.GetReportPartialTime().
//
//     -Reloads word filters with filters.LoadFiltersWord().
//
//   - If ctx is cancelled, the function returns.
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
