package exporter

import (
	"net"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// ShipLogs sends accumulated logs to a remote server (shipper) using the configured
// protocol and address.
//
// Parameters:
//
//   - logs: map containing log statistics grouped by source and message.
//
// Returns:
//
//   - err: error object if the connection, data transformation, or transmission fails.
//
// The function performs:
//
//   - Checks if 'options.Config.ShipperConfig.Address' is empty; if so, returns nil.
//
//   - Establishes a network connection using the configured protocol, address, and
//     timeout (based on FlushInterval).
//
//   - Uses 'defer' to ensure the connection is closed after the operation.
//
//   - Calls ShipsIntegrations(logs) to transform the log map into a JSON-encoded byte slice.
//
//   - Writes the resulting byte slice to the established connection.
//
//   - Returns the error if any step (dialing, transforming, or writing) fails;
//     otherwise, returns nil.
func ShipLogs(logs map[string]map[string]*pipelines.LogStats) (err error) {
	if options.Config.ShipperConfig.Address == "" {
		return nil
	}
	conn, err := net.DialTimeout(
		options.Config.ShipperConfig.Protocol,
		options.Config.ShipperConfig.Address,
		time.Duration(options.Config.ShipperConfig.FlushInterval)*time.Second,
	)

	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := ShipsIntegrations(logs)
	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	return nil
}
