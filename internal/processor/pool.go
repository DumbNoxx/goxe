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
	logs      = make(map[string]map[string]*pipelines.LogStats, 100)
	logsBurst = make(map[string]*pipelines.LogBurst, 100)
	errs      = []string{
		"ERROR",
		"CRITICAL",
	}
)

// Main function that processes the received information and sends it to their corresponding functions
func Clean(pipe <-chan *pipelines.LogEntry, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	ticker := time.NewTicker(60 * time.Minute)
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
			buf := text.RawEntry
			sanitizadedText = cluster.Cluster(text.Content, text.IdLog)
			if len(sanitizadedText) < 3 {
				text.Content = ""
				text.IdLog = ""
				text.Source = ""
				text.Timestamp = time.Time{}
				text.RawEntry = nil
				pipelines.EntryPool.Put(text)
				pipelines.BufferPool.Put(buf)
				continue
			}
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
			burstDetection(logsBurst, word)
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
			logsToFlush := logs
			exporter.Console(logsToFlush, mu, false)
			err := exporter.ShipLogs(logsToFlush)
			if err != nil {
				log.Print("Error sent")
			}
			clear(logs)
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
	if slices.Contains(errs, word) {
		stats := logsBurst[word]
		elapsed := time.Since(logsBurst[word].WindowStart)
		stats.Count++
		if elapsed > limitBreak {
			logsBurst[word].WindowStart = time.Now()
			logsBurst[word].Count = 1
			logsBurst[word].AlertsSent = 0
			logsBurst[word].LastAlertTime = time.Time{}
			goto CheckGlobal
		}

		if stats.Count <= 10 {
			goto CheckGlobal
		}

		if stats.AlertsSent >= 10 || time.Since(logsBurst[word].LastAlertTime) < 5*time.Second {
			return
		}

		handleWebhook(word, logsBurst[word])
		logsBurst[word].LastAlertTime = time.Now()
		logsBurst[word].AlertsSent++

		return
	}
CheckGlobal:
	global, ok := logsBurst["AGGREGATE_TRAFFIC"]

	if !ok {
		global = &pipelines.LogBurst{
			Count:         0,
			Category:      "AGGREGATE_TRAFFIC",
			WindowStart:   time.Now(),
			LastAlertTime: time.Time{},
			AlertsSent:    0,
		}
		logsBurst["AGGREGATE_TRAFFIC"] = global
	}

	global.Count++
	elapsedGlobal := time.Since(global.WindowStart)

	if elapsedGlobal > limitBreak {
		global.Count = 0
		global.WindowStart = time.Now()
		global.AlertsSent = 1
		global.LastAlertTime = time.Time{}
		return
	}

	if global.Count < 100 {
		return
	}

	if global.AlertsSent >= 10 || time.Since(global.LastAlertTime) < 5*time.Second {
		return
	}

	handleWebhook("AGGREGATE_TRAFFIC", global)
	global.LastAlertTime = time.Now()
	global.AlertsSent++
}

func handleWebhook(msg string, stats *pipelines.LogBurst) {
	var (
		data []byte
		err  error
	)
	for _, url := range options.Config.WebHookUrls {
		if strings.HasPrefix(url, "https://discord.com") {
			var DataSentWebhook pkg.WebhookDiscord
			var log = pkg.OptionsEmbedsDiscord{
				Title:       msg,
				Description: "The server's acting up.",
				Color:       16777215,
				Author: pkg.AuthorOptionsEmbedsDiscord{
					Name:    "Goxe",
					Url:     "https://github.com/DumbNoxx/Goxe",
					IconUrl: "https://raw.githubusercontent.com/DumbNoxx/Dotfiles-For-Humans/refs/heads/main/src/assets/img/goxe.png",
				},
				Fields: []pkg.FieldEmbedsDiscord{
					{
						Name:   "Errors",
						Value:  "```Check the server, it's overheating.```",
						Inline: false,
					},
					{
						Name:   "Category",
						Value:  stats.Category,
						Inline: true,
					},
					{
						Name:   "Start Time",
						Value:  stats.WindowStart.Format("02-01-2006, 15:04"),
						Inline: true,
					},
					{
						Name:   "Counts",
						Value:  fmt.Sprintf("%d", stats.Count),
						Inline: true,
					},
				},
				Footer: pkg.FooterEmbedsDiscord{
					Text: "Your Log Collector ❤️",
				},
				Timestamp: time.Now(),
			}
			DataSentWebhook.Embeds = append(DataSentWebhook.Embeds, log)
			data, err = json.Marshal(DataSentWebhook)
			sentData(data, err, url)
		}

		if strings.HasPrefix(url, "https://hooks.slack.com") {
			var headerLog = pkg.OptionsBlockSlack{
				Type: "header",
				Text: &pkg.OptionsTextMrkSlack{
					Type:  "plain_text",
					Text:  msg,
					Emoji: true,
				},
			}

			var mrkLog = pkg.OptionsBlockSlack{
				Type: "section",
				Text: &pkg.OptionsTextMrkSlack{
					Type: "mrkdwn",
					Text: fmt.Sprintf(
						"```Check the server, it's overheating.\nCount: %d - Start Time: %v - Category: %s```",
						stats.Count,
						stats.WindowStart.Format("02-01-2006, 15:04"),
						stats.Category,
					),
				},
			}

			var divider = pkg.OptionsBlockSlack{
				Type: "divider",
			}

			var footerLog = pkg.OptionsBlockSlack{
				Type: "context",
				Elements: []pkg.OptionsElementsBlockSlack{
					{
						Type:  "plain_text",
						Text:  "Author: Goxe",
						Emoji: true,
					},
				},
			}

			payload := pkg.WebhookSlack{
				Blocks: []pkg.OptionsBlockSlack{
					headerLog,
					mrkLog,
					divider,
					footerLog,
				},
			}

			data, err = json.Marshal(payload)
			sentData(data, err, url)
		}
	}
}

func sentData(data []byte, err error, url string) {
	options.SentWebhook(url, data)
	if err != nil {
		log.Print("Convert json fail")
		return
	}
}
