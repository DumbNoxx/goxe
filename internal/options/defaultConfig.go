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
		GenerateLogsOptions: options.GenerateLogsOptions{
			GenerateLogsFile: false,
			Hour:             "00:00:00",
		},
		WebHookUrls: []string{},
		BurstDetectionOptions: options.BurstDetectionOptions{
			LimitBreak: 10,
		},
		ShipperConfig: options.ShipperConfig{
			Address:       "",
			FlushInterval: 30,
			Protocol:      "tcp",
		},
		BufferUdpSize: 4,
	}
}
