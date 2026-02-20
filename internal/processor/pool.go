package processor

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"unsafe"

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
	logs             = make(map[string]map[string]*pipelines.LogStats, 100)
	logsBurst        = make(map[string]*pipelines.LogBurst, 100)
	TimeReport       time.Duration
	logsToFile       = make([]map[string]map[string]*pipelines.LogStats, 0)
	TickerReportFile *time.Ticker
	Ticker           *time.Ticker
)

// init loads the word filters and configures the tickers for periodic log
// exporting. It executes automatically upon package import.
//
// - TickerReportFile: controls report file generation.
// - Ticker: controls periodic export (console/remote) based on ReportInterval.
func init() {
	filters.LoadFiltersWord()
	TickerReportFile = time.NewTicker(utils.TimeReportFile)
	TimeReport = time.Duration(options.Config.ReportInterval * float64(time.Minute))
	Ticker = time.NewTicker(TimeReport)
}

// Clean function that processes the received information and sends it to their corresponding functions
//
// Parameters:
//   - ctx: context for cancellation (not explicitly used in the current body)
//   - pipe: input channel of *pipelines.LogEntry. Processing continues until
//     the channel is closed and all remaining entries are drained.
//   - wg: WaitGroup to notify the caller that the goroutine has completed
//   - mu: mutex protecting the shared maps 'logs' and 'logsBurst'.
//
// The function performs:
//   - Sanitization and clustering of log content.
//   - Statistics updates by source and message (logs).
//   - Burst detection by log level (logsBurst).
//   - Returning processed objects back to pools (EntryPool, BufferPool).
//   - Periodic exporting triggered by Ticker.C
//     Sends logs to exporter.ShipLogs and exporter.Console.
//     If GenerateLogsFile is enabled, accumulates logs for later writing.
//   - File Export triggered by TickerReportFile.C when GenerateLogsFile is true.
//
// Note: This functions is intented to run as a concurrent goroutine.
// It uses the unsafe package for zero-copy byte-to-string conversions,
// assuming the underlying buffers will not be modified afterward.
func Clean(ctx context.Context, pipe <-chan *pipelines.LogEntry, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	defer Ticker.Stop()
	defer TickerReportFile.Stop()

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
				exporter.ShipLogs(logs)
				return
			}
			buf := text.RawEntry
			dataCluster := cluster.Cluster(text.Content, text.IdLog)
			sanitizadedText = unsafe.String(unsafe.SliceData(dataCluster), len(dataCluster))
			mu.Lock()
			if logs[text.Source] == nil {
				logs[text.Source] = make(map[string]*pipelines.LogStats)
			}
			sliceData := sanitizer.ExtractLevelUpper(text.Content)
			word := unsafe.String(unsafe.SliceData(sliceData), len(sliceData))
			if logs[text.Source][sanitizadedText] == nil {
				logs[text.Source][sanitizadedText] = &pipelines.LogStats{
					Count:     0,
					FirstSeen: text.Timestamp,
					LastSeen:  text.Timestamp,
					Level:     []byte(word),
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
			text.Content = []byte("")
			text.IdLog = ""
			text.Source = ""
			text.Timestamp = time.Time{}
			text.RawEntry = nil
			pipelines.EntryPool.Put(text)
			pipelines.BufferPool.Put(buf)
		case <-Ticker.C:
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
		case <-TickerReportFile.C:
			if !options.Config.GenerateLogsOptions.GenerateLogsFile {
				continue
			}

			mu.Lock()
			logsToFlush := logsToFile
			logsToFile = make([]map[string]map[string]*pipelines.LogStats, 0)
			mu.Unlock()
			exporter.File(logsToFlush)
			exporter.ShipLogsFile(logsToFlush)
		}
	}
}
