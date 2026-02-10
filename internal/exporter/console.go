package exporter

import (
	"fmt"
	"strings"

	"github.com/DumbNoxx/goxe/internal/utils/colors"
	pipelines "github.com/DumbNoxx/goxe/pkg/pipelines"
)

// This function receives the map of logs created by the processor
func Console(logs map[string]map[string]*pipelines.LogStats, isFinal bool) {

	if isFinal {
		fmt.Println(strings.ToUpper("\tFinal Report"))
	} else {
		fmt.Println(strings.ToUpper("\tPartial Report"))
	}
	fmt.Println("----------------------------------")

	if len(logs) == 0 {
		return
	}

	for key, messages := range logs {

		fmt.Printf("ORIGIN: [%s]\n", key)

		if len(messages) == 0 {
			continue
		}

		for msg, stats := range messages {
			fmt.Printf("-%s [%d] %s %s -- (First seen %v - Last seen %v)\n", colors.GREEN, stats.Count, msg, colors.RESET, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
		}
	}

	fmt.Println("----------------------------------")
	memoryUsage()
}
