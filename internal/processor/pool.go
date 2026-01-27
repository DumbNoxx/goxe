package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/exporter"
	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/internal/processor/cluster"
	"github.com/DumbNoxx/Goxe/internal/processor/sanitizer"
	"github.com/DumbNoxx/Goxe/internal/utils"
	pkg "github.com/DumbNoxx/Goxe/pkg/options"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	logs      = make(map[string]map[string]*pipelines.LogStats)
	logsBurst = make(map[string]*pipelines.LogBurst)
)

var errs = []string{
	"ERROR",
	"CRITICAL",
}

// Main function that processes the received information and sends it to their corresponding functions
func Clean(pipe <-chan *pipelines.LogEntry, wg *sync.WaitGroup, mu *sync.Mutex) {
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
			word := sanitizer.ExtractLevelUpper(text.Content)
			if logsBurst[word] == nil {
				logsBurst[word] = &pipelines.LogBurst{
					Count:       0,
					Category:    word,
					WindowStart: time.Now(),
				}
			}
			burstDetection(logsBurst, word)
			logs[text.Source][sanitizadedText].Count++
			logs[text.Source][sanitizadedText].LastSeen = text.Timestamp
			mu.Unlock()
		case <-ticker.C:
			if len(logs) <= 0 {
				continue
			}
			exporter.Console(logs, mu, false)
		case <-tickerReportFile.C:
			if !options.Config.GenerateLogsOptions.GenerateLogsFile {
				continue
			}

			mu.Lock()
			logsToFlush := logs
			clear(logs)
			mu.Unlock()
			exporter.File(logsToFlush)
		}
	}

}

func burstDetection(logsBurst map[string]*pipelines.LogBurst, word string) {
	limitBreak := time.Second * time.Duration(options.Config.BurstDetectionOptions.LimitBreak)
	global, ok := logsBurst["AGGREGATE_TRAFFIC"]

	if !ok {
		global = &pipelines.LogBurst{
			Count:       0,
			Category:    "AGGREGATE_TRAFFIC",
			WindowStart: time.Now(),
		}
		logsBurst["AGGREGATE_TRAFFIC"] = global
	}

	global.Count++
	elapsedGlobal := time.Since(global.WindowStart)
	if global.Count > 100 && elapsedGlobal <= limitBreak {
		handleWebhook("DDos detected")
	}
	if elapsedGlobal > limitBreak {
		global.Count = 1
		global.WindowStart = time.Now()
	}

	if !slices.Contains(errs, word) {
		return
	}

	elapsed := time.Since(logsBurst[word].WindowStart)
	logsBurst[word].Count++
	if logsBurst[word].Count > 10 && elapsed <= limitBreak {
		handleWebhook("Critical System Errors")
	}
	if elapsed > limitBreak {
		logsBurst[word].WindowStart = time.Now()
		logsBurst[word].Count = 1
	}
}

func handleWebhook(text string) {
	var (
		data []byte
		err  error
	)

	for _, url := range options.Config.WebHookUrls {

		if strings.HasPrefix(url, "https://discord.com") {
			message := pkg.WebhookDiscord{
				Content: text,
			}
			data, err = json.Marshal(message)
			sentData(data, err, url)
		}

		if strings.HasPrefix(url, "https://hooks.slack.com") {
			message := pkg.WebhookSlack{
				Text: text,
			}
			data, err = json.Marshal(message)
			sentData(data, err, url)
		}
	}
}

func sentData(data []byte, err error, url string) {
	options.SentWebhook(url, []byte(data))
	if err != nil {
		log.Print("Convert json fail")
		return
	}
}
