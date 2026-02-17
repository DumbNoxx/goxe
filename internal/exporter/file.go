package exporter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// File writes a log report to a file within the user's cache directory.
//
// Parameters:
//
//   - logs: a slice of log maps (each map follows the structure: source -> message -> statistics).
//
// Returns:
//
//   - void: no return value; however, it logs warnings via log.Printf or terminates the program with log.Fatal upon critical failure.
//
// The function performs:
//
//   - Retrieves the user's cache directory using os.UserCacheDir(). If it fails, a warning is printed and the process continues.
//
//   - Generates a filename based on the current date: "log_YYYY-MM-DD.log".
//
//   - Constructs the full path: <cacheDir>/goxe/logs/<filename>.
//
//   - Attempts to read the file location with os.ReadDir. If an error occurs that is not "not found," it terminates with log.Fatal.
//
//   - Builds the report content using a strings.Builder:
//
//     -Includes a line with the configured hour (options.Config.GenerateLogsOptions.Hour).
//
//     -Adds separators for clarity.
//
//     -For each map in logs, and for each source and message, it writes the origin, counter, message, and timestamps.
//
//     -Adds a final separator.
//
//   - If the file does not exist (os.IsNotExist(err)), it writes the content using os.WriteFile with 0600 permissions.
//
// Note:
//   - The function only creates the file if it does not already exist; it does not overwrite or append to existing files.
//   - It does not create the "logs" directory if missing; it assumes the directory was previously created (e.g., by CacheDirGenerate).
func File(logs []map[string]map[string]*pipelines.LogStats) {
	cacheDir, cacheDirErr := os.UserCacheDir()
	if cacheDirErr != nil {
		log.Printf("Could not determine cache directory: %v. Using default settings based on: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html", cacheDirErr)
	}

	date := time.Now().Format("2006-01-02")

	var (
		folderCachePath string
		err             error
		data            strings.Builder
	)

	file := fmt.Sprintf("log_%s.log", date)

	folderCachePath = filepath.Join(cacheDir, "goxe", "logs", file)
	_, err = os.ReadDir(folderCachePath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	fmt.Fprintf(&data, "DIARY REPORT - Set time: [%v]\n", options.Config.GenerateLogsOptions.Hour)

	fmt.Fprintln(&data, "----------------------------------")
	for _, messages := range logs {

		if len(messages) == 0 {
			continue
		}

		for key, stat := range messages {
			fmt.Fprintf(&data, "ORIGIN: [%s]\n", key)
			for msg, stats := range stat {
				fmt.Fprintf(&data, "- [%d] %s -- (First seen %v - Last seen %v)\n", stats.Count, msg, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
			}
		}
	}

	fmt.Fprintln(&data, "----------------------------------")

	if os.IsNotExist(err) {
		err = os.WriteFile(folderCachePath, []byte(data.String()), 0600)
	}

}
