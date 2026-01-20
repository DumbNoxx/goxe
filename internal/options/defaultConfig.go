package options

import "github.com/DumbNoxx/Goxe/pkg/options"

func configDefault() options.Config {

	return options.Config{
		Port:          1729,
		IdLog:         "",
		PatternsWords: []string{},
		GenerateLogsOptions: options.OptionsGenerateLogs{
			GenerateLogs: false,
			Hour:         "00:00:00",
		},
	}
}
