package exporter

import (
	"fmt"
	"strings"

	"github.com/DumbNoxx/goxe/internal/utils/colors"
	pipelines "github.com/DumbNoxx/goxe/pkg/pipelines"
)

// Console prints a formatted and color-coded log report to the terminal.
//
// Parameters:
//
//   - logs: a map of maps where the outer key is the source and the inner map contains messages and their statistics (*pipelines.LogStats).
//   - isFinal: indicates whether this is the final report (true) or a partial one (false); affects the displayed title.
//
// Returns:
//
//   - void: no return value; the function's primary purpose is standard output.
//
// The function performs:
//
//   - Prints a "Final Report" or "Partial Report" title based on isFinal, in uppercase and with indentation.
//
//   - Prints a separator line: "----------------------------------".
//
//   - If the logs map is empty, returns without printing additional content.
//
//   - For each source in logs:
//
//     -Prints "ORIGIN: [source]".
//
//     -If the source contains no messages, skips to the next one.
//
//     -For each message in that source, prints a green-colored line containing the counter, the message,
//     and the timestamps for first and last seen (format: "15:04:05").
//
//   - At the end, prints another separator line and calls memoryUsage() to display the current memory footprint.
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
