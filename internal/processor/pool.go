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

func init() {
	filters.LoadFiltersWord()
	TickerReportFile = time.NewTicker(utils.TimeReportFile)
	TimeReport = time.Duration(options.Config.ReportInterval * float64(time.Minute))
	Ticker = time.NewTicker(TimeReport)
}

// Main function that processes the received information and sends it to their corresponding functions
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
		}
	}
}
