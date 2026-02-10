package burstdetection

import (
	"slices"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	webhooks "github.com/DumbNoxx/goxe/internal/processor/burstDetection/Webhooks"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	errs = []string{
		"ERROR",
		"CRITICAL",
	}
)

func BurstDetection(logsBurst map[string]*pipelines.LogBurst, word string) {
	limitBreak := time.Duration(float64(time.Second) * options.Config.BurstDetectionOptions.LimitBreak)
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

		webhooks.HandleWebhook(word, logsBurst[word])
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

	webhooks.HandleWebhook("AGGREGATE_TRAFFIC", global)
	global.LastAlertTime = time.Now()
	global.AlertsSent++
}
