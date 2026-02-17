package logs

import (
	"log"
	"os"
	"path/filepath"
)

// LogsCacheDirGenerate creates a "logs" directory within the specified path
// to store log cache.
//
// Parameters:
//
//   - folderCachePath: base path where the "logs" subfolder will be created.
//
// Returns:
//
//   - void: returns no value, but logs errors using log.Printf if the creation fails.
//
// The function performs:
//
//   - Builds the full path by concatenating folderCachePath and "logs".
//   - Calls os.MkdirAll with 0700 permissions to create the directory and its parents if necessary.
//   - If an error occurs, it prints a message with log.Printf indicating the path and the error.
func LogsCacheDirGenerate(folderCachePath string) {
	logsCachePath := filepath.Join(folderCachePath, "logs")
	err := os.MkdirAll(logsCachePath, 0700)
	if err != nil {
		log.Printf("Error create folder in %v, error: %v", logsCachePath, err)
	}
}
