package processor

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
	"unsafe"

	"github.com/DumbNoxx/goxe/internal/exporter"
	"github.com/DumbNoxx/goxe/internal/processor/cluster"
	"github.com/DumbNoxx/goxe/internal/processor/integrations"
	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// CleanFile processes a log file line by line, applying sanitization and clustering,
// accumulating statistics, and sending them to exporters (ShipLogs and FileReader).
//
// Parameters:
//
//   - file: pointer to the opened file to be read.
//   - idLog: log identifier (passed to cluster.Cluster).
//   - mu: mutex (maintained for consistency with other signatures, though logsFile is local).
//   - routeFile: path of the original file (passed to exporter.FileReader for reference).
//
// The function performs:
//
//   - Initializes a local map 'logsFile' to accumulate statistics under the "file-reader" source.
//
//   - Creates a bufio.Scanner to read the file line by line.
//
//   - For each line:
//
//     -Reads line bytes.
//
//     -Calls cluster.Cluster(data, idLog) to obtain sanitized and grouped text.
//
//     -Converts the result to a string using 'unsafe' (zero-copy) and assigns it to sanitizedText.
//
//     -Acquires the mutex (kept for structural consistency).
//
//     -Initializes the entry in logsFile["file-reader"] if necessary.
//
//     -Extracts the log level using sanitizer.ExtractLevelUpper(data) and converts it to string via 'unsafe'.
//
//     -If the sanitized message is new, creates an entry in logsFile with initial statistics.
//
//     -Increments the counter and updates LastSeen.
//
//     -Releases the mutex.
//
//   - After reaching EOF, calls exporter.ShipLogs(logsFile) to send logs to the remote shipper.
//
//   - Calls exporter.FileReader(logsFile, routeFile) to save the report to a file.
//
//   - Calls integrations.Integrations(logs), send data to Observability Platforms
//
//   - Checks for scanner errors via scanner.Err() and terminates with log.Fatal if found.
//
//   - Clears the logsFile map using 'clear()' to free memory.
func CleanFile(file *os.File, idLog string, mu *sync.Mutex, routeFile string) {
	var (
		sanitizadedText string
		data            []byte
		logsFile        = make(map[string]map[string]*pipelines.LogStats, 100)
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = scanner.Bytes()
		dataCluster := cluster.Cluster(data, idLog)
		sanitizadedText = unsafe.String(unsafe.SliceData(dataCluster), len(dataCluster))

		mu.Lock()
		if logsFile["file-reader"] == nil {
			logsFile["file-reader"] = make(map[string]*pipelines.LogStats)
		}
		sliceData := sanitizer.ExtractLevelUpper(data)
		word := unsafe.String(unsafe.SliceData(sliceData), len(sliceData))
		if logsFile["file-reader"][sanitizadedText] == nil {
			logsFile["file-reader"][sanitizadedText] = &pipelines.LogStats{
				Count:     0,
				FirstSeen: time.Now(),
				LastSeen:  time.Now(),
				Level:     []byte(word),
			}
		}
		logsFile["file-reader"][sanitizadedText].Count++
		logsFile["file-reader"][sanitizadedText].LastSeen = time.Now()
		mu.Unlock()
	}
	err := exporter.ShipLogs(logsFile)
	if err != nil {
		log.Fatal(err)
	}
	exporter.FileReader(logsFile, routeFile)
	integrations.Integrations(logsFile)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	clear(logsFile)
}
