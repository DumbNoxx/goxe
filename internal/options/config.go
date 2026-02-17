package options

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/DumbNoxx/goxe/pkg/options"
)

var Config = ConfigFile()

// ConfigFile loads the configuration from the config.json file in the user's config directory.
// If the file does not exist, it creates it with default values. If it exists but is invalid, it uses the default values.
//
// The function performs:
//
//   - Retrieves the user's configuration directory using os.UserConfigDir().
//
//   - If it fails, it prints a warning (continues with an empty dir, which will likely cause errors).
//
//   - Builds the full path: filepath.Join(dir, "goxe", "config.json").
//
//   - Attempts to read the file using os.ReadFile.
//
//   - If an error other than "does not exist" occurs, it terminates the program with log.Fatal.
//
//   - If the file does not exist:
//
//     -Creates the parent directory using os.MkdirAll (0700 permissions).
//
//     -Writes the default configuration (obtained from configDefault()) in indented JSON format with 0600 permissions.
//
//   - If writing fails, it prints an error with log.Printf.
//
//   - Uses the default data (bDefault) as the read content.
//
//   - Attempts to deserialize the content (origConfig) into the config variable using json.Unmarshal.
//
//   - If deserialization fails, it prints a warning and assigns the default configuration (configDefault).
//
//   - Returns the configuration (either read from the file or the default one).
func ConfigFile() (config options.Config) {
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
