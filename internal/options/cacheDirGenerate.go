package options

import (
	"log"
	"os"
	"path/filepath"

	"github.com/DumbNoxx/goxe/internal/options/logs"
)

// CacheDirGenerate creates the cache directory structure for the application.
//
// Returns:
//
//   - void: returns no value, but may terminate the program with log.Fatal if reading the directory fails.
//
// The function performs:
//
//   - Retrieves the user's cache directory using os.UserCacheDir().
//   - If it fails, it prints a warning message with log.Printf (continues with an empty dir, which will cause later errors).
//   - Builds the full path: filepath.Join(dir, "goxe").
//   - Attempts to read the directory using os.ReadDir.
//   - If an error other than "does not exist" occurs, it terminates the program with log.Fatal(err).
//   - If the directory does not exist (os.IsNotExist), it creates it using os.MkdirAll(folderCachePath, 0700).
//   - If the creation fails, it prints an error with log.Printf (does not terminate the program).
//   - If Config.GenerateLogsOptions.GenerateLogsFile is true, it calls logs.LogsCacheDirGenerate(folderCachePath)
//     to create the "logs" subdirectory.
func CacheDirGenerate() {
	dir, dirErr := os.UserCacheDir()
	if dirErr != nil {
		log.Printf("Could not determine cache directory: %v. Using default settings based on: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html", dirErr)
	}

	var (
		folderCachePath string
		err             error
	)

	folderCachePath = filepath.Join(dir, "goxe")
	_, err = os.ReadDir(folderCachePath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if os.IsNotExist(err) {
		err := os.MkdirAll(folderCachePath, 0700)
		if err != nil {
			log.Printf("Error create folder in %v, error: %v", folderCachePath, err)
		}
	}

	if Config.GenerateLogsOptions.GenerateLogsFile {
		logs.LogsCacheDirGenerate(folderCachePath)
	}
}
