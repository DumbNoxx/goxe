package exporter

import (
	"fmt"
	"sync"

	"github.com/DumbNoxx/Goxe/internal/utils/colors"
	logslevel "github.com/DumbNoxx/Goxe/internal/utils/logsLevel"
)

// This function receives the map of logs created by the processor
func Console(messages map[string]int, mu *sync.Mutex) {
	fmt.Println("\tPartial Report")
	println("")
	mu.Lock()
	for msg, count := range messages {
		switch {
		case count >= logslevel.CRITIC:
			fmt.Printf("%s%d Veces: %s%s\n", colors.RED, count, msg, colors.RESET)
		case count >= logslevel.NORMAL:
			fmt.Printf("%s%d Veces: %s%s\n", colors.YELLOW, count, msg, colors.RESET)
		case count <= logslevel.SAVED:
			fmt.Printf("%s%d Veces: %s%s\n", colors.GREEN, count, msg, colors.RESET)
		}
	}
	mu.Unlock()
}
