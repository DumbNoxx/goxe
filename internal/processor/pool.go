package processor

import (
	"fmt"
	"sync"
	"time"
)

var (
	messages = make(map[string]int)
	mu       sync.Mutex
)

func Clean(pipe chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case text, ok := <-pipe:
			if !ok {
				return
			}
			mu.Lock()
			messages[text]++
			mu.Unlock()
		case <-ticker.C:
			fmt.Println("\t////Reporte parcial\\\\\\\\")
			mu.Lock()
			for msg, count := range messages {
				fmt.Printf("%d veces: %s\n", count, msg)
			}
			mu.Unlock()
		}
	}

}
