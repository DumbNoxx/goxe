package options

import (
	"log"
	"os"
	"path/filepath"

	"github.com/DumbNoxx/Goxe/internal/options/logs"
)

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

	if Config.GenerateLogsOptions.GenerateLogs {
		logs.LogsCacheDirGenerate(folderCachePath)
	}
}
