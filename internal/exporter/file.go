package exporter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

func File(logs map[string]map[string]*pipelines.LogStats) {
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

	fmt.Fprintf(&data, "PARTIAL REPORT - Set time: [%v]\n", options.Config.GenerateLogsOptions.Hour)

	fmt.Fprintln(&data, "----------------------------------")
	for key, messages := range logs {

		fmt.Fprintf(&data, "ORIGIN: [%s]\n", key)

		if len(messages) == 0 {
			continue
		}

		for msg, stats := range messages {
			fmt.Fprintf(&data, "- [%d] %s -- (First seen %v - Last seen %v)\n", stats.Count, msg, stats.FirstSeen.Format("15:04:05"), stats.LastSeen.Format("15:04:05"))
		}
	}

	fmt.Fprintln(&data, "----------------------------------")

	if os.IsNotExist(err) {
		err = os.WriteFile(folderCachePath, []byte(data.String()), 0600)
	}

}
