package processor

import (
	"fmt"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/exporter"
	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/internal/processor/cluster"
	"github.com/DumbNoxx/Goxe/internal/utils"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	logs = make(map[string]map[string]*pipelines.LogStats)
)

// Main function that processes the received information and sends it to their corresponding functions
func Clean(pipe <-chan pipelines.LogEntry, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	tickerReportFile := time.NewTicker(utils.UserConfigHour())
	defer tickerReportFile.Stop()

	var sanitizadedText string

	for {
		select {
		case text, ok := <-pipe:
			if !ok {
				if len(logs) <= 0 {
					fmt.Println("\n[System] System terminated")
					return
				}
				fmt.Println("\n[System] System terminated last report")
				exporter.Console(logs, mu, true)
				return
			}
			sanitizadedText = cluster.Cluster(text.Content, text.IdLog)
			if len(sanitizadedText) < 3 {
				continue
			}
			mu.Lock()
			if logs[text.Source] == nil {
				logs[text.Source] = make(map[string]*pipelines.LogStats)
			}
			if logs[text.Source][sanitizadedText] == nil {

				logs[text.Source][sanitizadedText] = &pipelines.LogStats{
					Count:     0,
					FirstSeen: text.Timestamp,
					LastSeen:  text.Timestamp,
					Level:     text.Level,
				}

			}
			logs[text.Source][sanitizadedText].Count++
			logs[text.Source][sanitizadedText].LastSeen = text.Timestamp
			mu.Unlock()
		case <-ticker.C:
			if len(logs) <= 0 {
				continue
			}
			exporter.Console(logs, mu, false)
		case <-tickerReportFile.C:
			if !options.Config.GenerateLogsOptions.GenerateLogs {
				continue
			}

			mu.Lock()
			logsToFlush := logs
			logs = make(map[string]map[string]*pipelines.LogStats)
			mu.Unlock()
			exporter.File(logsToFlush)
		}
	}

}
