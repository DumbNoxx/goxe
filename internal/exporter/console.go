package exporter

import (
	"fmt"
	"strings"
	"sync"

	"github.com/DumbNoxx/Goxe/internal/pipelines"
	"github.com/DumbNoxx/Goxe/internal/utils/colors"
	logslevel "github.com/DumbNoxx/Goxe/internal/utils/logsLevel"
)

// This function receives the map of logs created by the processor
func Console(logs map[string]map[string]*pipelines.LogStats, mu *sync.Mutex) {

	fmt.Println(strings.ToUpper("\tPartial Report"))
	fmt.Println("----------------------------------")

	mu.Lock()

	for key, messages := range logs {

		fmt.Printf("ORIGEN: [%s]\n", key)

		for msg, stats := range messages {

			switch {
			case stats.Count >= logslevel.CRITIC:
				fmt.Printf("- %s%s (x%d)%s -- (Last seen %v)\n", colors.RED, msg, stats.Count, colors.RESET, stats.LastSeen.Format("15:04:05"))
			case stats.Count >= logslevel.NORMAL:
				fmt.Printf("- %s%s (x%d)%s -- (Last seen %v)\n", colors.YELLOW, msg, stats.Count, colors.RESET, stats.LastSeen.Format("15:04:05"))
			case stats.Count <= logslevel.SAVED:
				fmt.Printf("- %s%s (x%d)%s -- (Last seen %v)\n", colors.GREEN, msg, stats.Count, colors.RESET, stats.LastSeen.Format("15:04:05"))
			}

		}
	}

	fmt.Println("----------------------------------")
	mu.Unlock()
}
