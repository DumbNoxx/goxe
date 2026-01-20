package logs

import (
	"log"
	"os"
	"path/filepath"
)

func LogsCacheDirGenerate(folderCachePath string) {
	logsCachePath := filepath.Join(folderCachePath, "logs")
	err := os.MkdirAll(logsCachePath, 0700)
	if err != nil {
		log.Printf("Error create folder in %v, error: %v", logsCachePath, err)
	}
}
