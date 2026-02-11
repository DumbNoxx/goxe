package processor

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/internal/exporter"
	"github.com/DumbNoxx/goxe/internal/options"
	burstdetection "github.com/DumbNoxx/goxe/internal/processor/burstDetection"
	"github.com/DumbNoxx/goxe/internal/processor/cluster"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
	"github.com/DumbNoxx/goxe/internal/utils"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	logs       = make(map[string]map[string]*pipelines.LogStats, 100)
	logsBurst  = make(map[string]*pipelines.LogBurst, 100)
	timeReport = time.Duration(options.Config.ReportInterval * float64(time.Minute))
	logsToFile = make([]map[string]map[string]*pipelines.LogStats, 0)
)

func init() {
	filters.LoadFiltersWord()
}

// Main function that processes the received information and sends it to their corresponding functions
func Clean(ctx context.Context, pipe <-chan *pipelines.LogEntry, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	ticker := time.NewTicker(timeReport)
	defer ticker.Stop()
	tickerReportFile := time.NewTicker(utils.TimeReportFile)
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
				exporter.Console(logs, true)
				return
			}
			buf := text.RawEntry
			sanitizadedText = cluster.Cluster(filters.Str.Replace(text.Content), text.IdLog)
			mu.Lock()
			if logs[text.Source] == nil {
				logs[text.Source] = make(map[string]*pipelines.LogStats)
			}
			word := sanitizer.ExtractLevelUpper(text.Content)
			if logs[text.Source][sanitizadedText] == nil {
				logs[text.Source][sanitizadedText] = &pipelines.LogStats{
					Count:     0,
					FirstSeen: text.Timestamp,
					LastSeen:  text.Timestamp,
					Level:     word,
				}
			}
			if logsBurst[word] == nil {
				logsBurst[word] = &pipelines.LogBurst{
					Count:         0,
					Category:      word,
					WindowStart:   time.Now(),
					AlertsSent:    0,
					LastAlertTime: time.Time{},
				}
			}
			burstdetection.BurstDetection(logsBurst, word)
			logs[text.Source][sanitizadedText].Count++
			logs[text.Source][sanitizadedText].LastSeen = text.Timestamp
			mu.Unlock()
			text.Content = ""
			text.IdLog = ""
			text.Source = ""
			text.Timestamp = time.Time{}
			text.RawEntry = nil
			pipelines.EntryPool.Put(text)
			pipelines.BufferPool.Put(buf)
		case <-ticker.C:

			if len(logs) <= 0 {
				continue
			}

			mu.Lock()
			logsToFlush := logs
			logs = make(map[string]map[string]*pipelines.LogStats, 100)
			mu.Unlock()
			if options.Config.GenerateLogsOptions.GenerateLogsFile {
				logsToFile = append(logsToFile, logsToFlush)
			}
			exporter.Console(logsToFlush, false)
			err := exporter.ShipLogs(logsToFlush)
			if err != nil {
				log.Print("Error sent")
			}
		case <-tickerReportFile.C:
			if !options.Config.GenerateLogsOptions.GenerateLogsFile {
				continue
			}

			mu.Lock()
			logsToFlush := logsToFile
			logsToFile = make([]map[string]map[string]*pipelines.LogStats, 0)
			mu.Unlock()
			exporter.File(logsToFlush)
		}
	}
}
