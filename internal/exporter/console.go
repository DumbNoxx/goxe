package exporter

import (
	"fmt"
	"strings"
	"sync"

	"github.com/DumbNoxx/Goxe/internal/utils/colors"
	logslevel "github.com/DumbNoxx/Goxe/internal/utils/logsLevel"
	pipelines "github.com/DumbNoxx/Goxe/pkg/pipelines"
)

// This function receives the map of logs created by the processor
func Console(logs map[string]map[string]*pipelines.LogStats, mu *sync.Mutex, isFinal bool) {

	if isFinal {
		fmt.Println(strings.ToUpper("\tFinal Report"))
	} else {
		fmt.Println(strings.ToUpper("\tPartial Report"))
	}
	fmt.Println("----------------------------------")

	mu.Lock()
	defer mu.Unlock()

	if len(logs) == 0 {
		return
	}

	for key, messages := range logs {

		fmt.Printf("ORIGIN: [%s]\n", key)

		if len(messages) == 0 {
			continue
		}

		for msg, stats := range messages {
			switch {
			case stats.Count >= logslevel.CRITIC:
				fmt.Printf("-%s [%d] %s %s -- (First seen %v - Last seen %v)\n", colors.RED, stats.Count, msg, colors.RESET, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
			case stats.Count >= logslevel.NORMAL:
				fmt.Printf("-%s [%d] %s %s -- (First seen %v - Last seen %v)\n", colors.YELLOW, stats.Count, msg, colors.RESET, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
			case stats.Count >= logslevel.SAVED:
				fmt.Printf("-%s [%d] %s %s -- (First seen %v - Last seen %v)\n", colors.GREEN, stats.Count, msg, colors.RESET, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
			}
		}
	}

	fmt.Println("----------------------------------")
	memoryUsage()
}
