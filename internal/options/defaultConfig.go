package options

import (
	"os"

	"github.com/DumbNoxx/Goxe/pkg/options"
)

func configDefault() options.Config {
	home, _ := os.Hostname()

	return options.Config{
		Port:          1729,
		IdLog:         home,
		PatternsWords: []string{},
		GenerateLogsOptions: options.OptionsGenerateLogs{
			GenerateLogs: false,
			Hour:         "00:00:00",
		},
	}
}
