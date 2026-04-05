package ingestor

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/DumbNoxx/goxe/internal/exporter"
	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor/cluster"
	"github.com/DumbNoxx/goxe/internal/processor/integrations"
	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
	pkgEx "github.com/DumbNoxx/goxe/pkg/exporter"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	logsFile = make(map[string]map[string]*pipelines.LogStats, 100)
)

// File defines the contract for processing and normalizing log files.
// It allows swapping different processing logic while maintaining a consistent signature.
type File interface {
	// FileNormalized handles reading, clustering, and exporting data from a file.
	FileNormalized(file *os.File, idLog string, mu *sync.Mutex, routeFile string, Shipper pkgEx.Shipper, getConfig options.ConfigProvider)
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
func (f *NormalizedManager) FileNormalized(file *os.File, idLog string, mu *sync.Mutex, routeFile string, Shipper pkgEx.Shipper, getConfig options.ConfigProvider) {

	var (
		sanitizadedText string
		data            []byte
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = scanner.Bytes()
		dataCluster := cluster.Cluster(data, idLog, getConfig)
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
	err := exporter.ShipLogs(logsFile, Shipper, getConfig)
	if err != nil {
		log.Fatal(err)
	}
	exporter.FileReader(logsFile, routeFile)
	integrations.Integrations(logsFile, Shipper, getConfig)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	clear(logsFile)
}

// JsonManager provides the specialized implementation for managing logs in JSON format.
// It focuses on structural analysis and recursive normalization of JSON data.
type JsonManager struct{}


// FileNormalized processes a JSON log file using a decoder to handle individual 
// objects or arrays, generating statistics via recursive analysis.
//
// Parameters:
//
//    - file: pointer to the opened JSON file to be decoded.
//    - idLog: log identifier (maintained for signature consistency).
//    - mu: mutex used to synchronize access to the global logsFile map.
//    - routeFile: path of the original file (passed to exporter.FileReaderJson).
//    - Shipper: instance of pkgEx.Shipper (maintained for interface compliance).
//
// The function performs:
//
//    - Initializes a json.Decoder to process the file stream.
//    - Decodes the content into a generic 'dataMap' (supports map[string]any or []any).
//    - For each decoded element:
//        - Recursively analyzes the structure via f.analizeJson to produce a unique ID.
//        - Persists the statistics via f.saveLog.
//    - Upon reaching EOF, calls exporter.FileReaderJson to generate the report.
//    - Terminates with log.Fatal if a decoding error (other than EOF) occurs.
//    - Clears the logsFile map using 'clear()' to free memory.
func (f *JsonManager) FileNormalized(file *os.File, idLog string, mu *sync.Mutex, routeFile string, Shipper pkgEx.Shipper, getConfig options.ConfigProvider) {
	var (
		dataMap any
		id      string
	)
	dec := json.NewDecoder(file)
	for {
		if err := dec.Decode(&dataMap); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err.Error())
		}
		switch va := dataMap.(type) {
		case map[string]any:
			id = f.analizeJson(va)
			f.saveLog(id, mu)
		case []any:
			for _, item := range va {
				id = f.analizeJson(item)
				f.saveLog(id, mu)
			}
		}

	}
	exporter.FileReaderJson(logsFile, routeFile)
	clear(logsFile)
}

// saveLog updates the global statistics map in a thread-safe manner.
//
// Parameters:
//
//    - id: the unique structural identifier generated by analizeJson.
//    - mu: pointer to the mutex for safe map manipulation.
//
// The function performs:
//    - Validates that the ID is not empty.
//    - Locks the mutex to ensure atomicity.
//    - Initializes the "file-reader" source in logsFile if it doesn't exist.
//    - Creates or updates the LogStats entry, incrementing the occurrence counter.
func (f *JsonManager) saveLog(id string, mu *sync.Mutex) {
	if id == "" {
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if logsFile["file-reader"] == nil {
		logsFile["file-reader"] = make(map[string]*pipelines.LogStats)
	}
	if logsFile["file-reader"][id] == nil {
		logsFile["file-reader"][id] = &pipelines.LogStats{Count: 0}
	}
	logsFile["file-reader"][id].Count++
}

// analizeJson performs a recursive structural analysis of a JSON element 
// to create a "fingerprint" or ID based on keys and value types.
//
// Parameters:
//
//    - file: any type representing a decoded JSON fragment.
//
// The function performs:
//    - Identifies types (bool, num, string) and replaces them with generic placeholders.
//    - For strings, uses f.isId to detect and mask potential unique identifiers (hashes/UUIDs).
//    - For arrays, iterates and builds a indexed string of its elements' structures.
//    - For maps, sorts keys alphabetically and recursively builds a key:value_structure pipe-separated string.
//    - Returns a normalized string representing the "shape" of the JSON.
func (f *JsonManager) analizeJson(file any) string {

	switch value := file.(type) {
	case nil:
		fmt.Println("nil")
		return ""
	case bool:
		return "<BOOL>"
	case float64, int:
		return "<NUM>"
	case string:
		if f.isId(value) {
			return "<ID>"
		}
		return value
	case []any:
		var indexTotal strings.Builder
		for index, item := range value {
			interId := f.analizeJson(item)
			fmt.Fprintf(&indexTotal, "%d:%s", index, interId)
		}
		return indexTotal.String()
	case map[string]any:
		keys := make([]string, 0, len(value))
		for k := range value {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var idTotal strings.Builder
		for _, key := range keys {
			interId := f.analizeJson(value[key])
			fmt.Fprintf(&idTotal, "%s:%s||", key, interId)
		}
		return idTotal.String()
	}
	return ""
}

// isId determines if a string value should be treated as a unique identifier 
// (like a UUID, Hex hash, or long token) rather than a generic message.
//
// Parameters:
//
//    - value: the string to be evaluated.
//
// Returns true if the string:
//    - Does not contain spaces.
//    - Is longer than 20 characters or is a valid hexadecimal string.
func (f *JsonManager) isId(value string) bool {
	if strings.Contains(value, " ") {
		return false
	}
	if len(value) > 20 {
		return true
	}
	if _, err := hex.DecodeString(value); err != nil {
		return true
	}

	return false
}
