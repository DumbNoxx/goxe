package options

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/DumbNoxx/Goxe/pkg/options"
)

var Config = configFile()

func configFile() (config options.Config) {
	dir, dirErr := os.UserConfigDir()

	configDefault := configDefault()

	bDefault, _ := json.MarshalIndent(configDefault, "", "  ")

	var (
		configPath string
		origConfig []byte
		err        error
	)

	if dirErr != nil {
		log.Printf("Could not determine config directory: %v. Using default settings based on: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html", dirErr)
	}

	configPath = filepath.Join(dir, "goxe", "config.json")

	origConfig, err = os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if os.IsNotExist(err) {

		err := os.MkdirAll(filepath.Dir(configPath), 0700)
		if err == nil {
			err = os.WriteFile(configPath, bDefault, 0600)
		}
		if err != nil {
			log.Printf("Error saving config changes: %v", err)
		}
		origConfig = bDefault
	}

	errUnmarshal := json.Unmarshal(origConfig, &config)
	if errUnmarshal != nil {
		log.Printf("Warning: config.json is corrupt (%v). Using internal defaults. [web default]", errUnmarshal)
		config = configDefault
	}
	return config
}
