package processor

import (
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/exporter"
	"github.com/DumbNoxx/Goxe/internal/pipelines"
)

var (
	logs = make(map[string]map[string]*pipelines.LogStats)
	mu   sync.Mutex
)

// Main function that processes the received information and sends it to their corresponding functions
func Clean(pipe <-chan pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	var sanitizadedText string

	for {
		select {
		case text, ok := <-pipe:
			if !ok {
				return
			}
			sanitizadedText = Sanitizador(text.Content)
			if len(sanitizadedText) < 3 {
				continue
			}
			mu.Lock()
			if logs[text.Source] == nil {
				logs[text.Source] = make(map[string]*pipelines.LogStats)
			}
			if logs[text.Source][sanitizadedText] == nil {
				logs[text.Source][sanitizadedText] = &pipelines.LogStats{
					Count:    0,
					LastSeen: text.Timestamp,
					Level:    text.Level,
				}
			}
			logs[text.Source][sanitizadedText].Count++
			logs[text.Source][sanitizadedText].LastSeen = text.Timestamp
			mu.Unlock()
		case <-ticker.C:
			exporter.Console(logs, &mu)
		}
	}

}
