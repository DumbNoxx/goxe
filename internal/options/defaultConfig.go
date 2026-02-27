package options

import (
	"os"

	"github.com/DumbNoxx/goxe/pkg/options"
)

// configDefault generates and returns a default configuration for the application.
//
// Returns:
//   - options.Config: struct with default values for all fields.
//
// The function performs:
//
//   - Retrieves the hostname using os.Hostname() (ignores the error) and assigns it to IdLog.
//
//   - Defines default values:
//
//     -Port: 1729
//
//     -IdLog: hostname
//
//     -PatternsWords: empty slice
//
//     -GenerateLogsOptions.GenerateLogsFile: false
//
//     -GenerateLogsOptions.Hour: "00:00:00"
//
//     -WebHookUrls: empty slice
//
//     -BurstDetectionOptions.LimitBreak: 10
//
//     -ShipperConfig.Address: ""
//
//     -ShipperConfig.FlushInterval: 30
//
//     -ShipperConfig.Protocol: "tcp"
//
//     -ReportInterval: 60
//
//     -BufferUdpSize: 4
//
//     -Integrations: empty slice
//
//     -Destination: empty string (used to determine the concrete Shipper implementation).
//
//   - Returns the complete structure.
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
		ReportInterval: 60,
		BufferUdpSize:  4,
		Integrations:   []options.IntegrationsShipper{},
		Destination:    "",
	}
}
