package ingestor

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
	pkgEx "github.com/DumbNoxx/goxe/pkg/exporter"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// File defines the contract for processing and normalizing log files.
// It allows swapping different processing logic while maintaining a consistent signature.
type File interface {
	// FileNormalized handles reading, clustering, and exporting data from a file.
	FileNormalized(file *os.File, idLog string, mu *sync.Mutex, routeFile string, Shipper pkgEx.Shipper)
}

// NormalizedManager provides the standard implementation for managing normalized logs.
// It centralizes sanitization logic and handles data distribution to multiple platforms.
type NormalizedManager struct{}

// FileNormalized processes a log file line by line, applying sanitization and clustering,
// accumulating statistics, and sending them to exporters (ShipLogs and FileReader).
//
// Parameters:
//
// - file: pointer to the opened file to be read.
// - idLog: log identifier (passed to cluster.Cluster).
// - mu: mutex (maintained for consistency with other signatures, though logsFile is local).
// - routeFile: path of the original file (passed to exporter.FileReader for reference).
// - Shipper: instance of pkgEx.Shipper used for data transmission.
//
// The function performs:
//
//   - Initializes a local map 'logsFile' with an initial capacity of 100 to accumulate
//     statistics under the "file-reader" source.
//
// - Creates a bufio.Scanner to read the file line by line.
//
// - For each line:
//
//   - Reads line bytes via scanner.Bytes().
//
//   - Calls cluster.Cluster(data, idLog) to obtain sanitized and grouped text.
//
//   - Converts the result to a string using 'unsafe' (zero-copy) and assigns it to sanitizedText.
//
//   - Acquires the mutex (mu.Lock) to maintain data consistency in the map.
//
//   - Extracts the log level using sanitizer.ExtractLevelUpper(data) and converts it to string via 'unsafe'.
//
//   - If the sanitized message is new, creates an entry in logsFile with initial statistics.
//
//   - Increments the counter and updates LastSeen.
//
//   - Releases the mutex (mu.Unlock).
//
// - After reaching EOF:
//
//   - Calls exporter.ShipLogs(logsFile, Shipper) to send logs to the remote shipper.
//     Fails with log.Fatal if an error occurs.
//
//   - Calls exporter.FileReader(logsFile, routeFile) to save the report to a file.
//
//   - Calls integrations.Integrations(logsFile, Shipper) to send data to Observability Platforms.
//
// - Checks for scanner errors via scanner.Err() and terminates with log.Fatal if found.
//
// - Clears the logsFile map using 'clear()' to free memory.
func (f *NormalizedManager) FileNormalized(file *os.File, idLog string, mu *sync.Mutex, routeFile string, Shipper pkgEx.Shipper) {

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
	err := exporter.ShipLogs(logsFile, Shipper)
	if err != nil {
		log.Fatal(err)
	}
	exporter.FileReader(logsFile, routeFile)
	integrations.Integrations(logsFile, Shipper)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	clear(logsFile)
}
