package processor

import (
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/exporter"
)

var (
	messages = make(map[string]int)
	mu       sync.Mutex
)

// Main function that processes the received information and sends it to their corresponding functions
func Clean(pipe chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	var sanitizadedText string

	for {
		select {
		case text, ok := <-pipe:
			if !ok {
				return
			}
			sanitizadedText = Sanitizador(text)
			if len(sanitizadedText) < 3 {
				continue
			}
			mu.Lock()
			messages[Sanitizador(text)]++
			mu.Unlock()
		case <-ticker.C:
			exporter.Console(messages, &mu)
		}
	}

}
