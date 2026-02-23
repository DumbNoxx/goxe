package burstdetection

import (
	"slices"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	webhooks "github.com/DumbNoxx/goxe/internal/processor/burstDetection/Webhooks"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	errs = []string{
		"ERROR",
		"CRITICAL",
	}
)

// BurstDetection detects and handles bursts based on severty level and AGGREGATE_TRAFFIC.
//
// Parameters:
//   - logsBurst: map storing burst states by category (e.g., 'ERROR', 'AGGREGATE_TRAFFIC'),
//   - word: current log level (e.g., 'ERROR', 'CRITICAL'), previously extracted from content.
//
// Returns:
//   - void: the function operates through side effects, modifying logsBurst and triggering webhooks
//
// The function performs:
//
//  1. If 'word' is in the error list (errs), it handles specific burst logic for that level:
//
//     - Calculates elapsed time since the window started (WindowStart).
//
//     - If the limit is exceeded (LimitBreak), it resets the window and counter, then jumps to CheckGlobal.
//
//     - If the counter is <= 10, it jumps to CheckGlobal (not yet burst).
//
//     - If 10 alerts have already been sent or the las one was less than 5 seconds ago, it does nothing (rate limiting).
//
//     - Otherwise, if triggers a webhook via webhooks.HandleWebhook, updates LastAlertTime, and increments AlertsSent.
//
//  2. CheckGlobal: label that unifies the flow to handle AGGREGATE_TRAFFIC bursts:
//
//     - Retrieves or creates the 'AGGREGATE_TRAFFIC' entry logsBurst.
//
//     - Increments its counter and checks the elapsed time.
//
//     - If it exceeds LimitBreak, it resets the window and alert counter.
//
//     - if the global counter is < 100, it does nothing.
//
//     - if 10 global alerts have been sent or the last one was less than 5 seconds ago, it does nothing.
//
//     - Otherwise, it triggers a webhook for AGGREGATE_TRAFFIC and updates its metada.
func BurstDetection(logsBurst map[string]*pipelines.LogBurst, word string, text []byte) {
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
			Ip:            filters.GetIpBurstDetection(text),
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
