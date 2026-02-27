package exporter

import (
	"net"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/exporter"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// ShipLogsFile sends accumulated logs to a remote server (shipper) using the configured protocol.
//
// Parameters:
//
//   - logs: a slice of maps where the outer key is the source and the inner map contains messages and their statistics (LogStats).
//
//   - Shipper: an interface of type exporter.Shipper used to format the data before sending.
//
// Returns:
//
//   - error: nil if the transmission was successful or no address is configured; otherwise, returns connection, formatting, or write errors.
//
// The function performs:
//
//   - If options.Config.ShipperConfig.Address is empty, it returns nil and does nothing.
//
//   - Establishes a connection using the protocol, address, and timeout specified in the configuration.
//
//   - Iterates through the slice of log maps:
//
//     -Calls Shipper.PrepareShip(messages) to transform each map of log statistics into a formatted byte slice.
//
//     -Writes the resulting byte slice to the established network connection.
//
//   - If any error occurs (connection, formatting via PrepareShip, or write), it returns immediately.
//
//   - Upon completion, closes the connection and returns nil.
func ShipLogsFile(logs []map[string]map[string]*pipelines.LogStats, Shipper exporter.Shipper) (err error) {
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

	for _, messages := range logs {
		data, err := Shipper.PrepareShip(messages)
		if err != nil {
			return err
		}
		_, err = conn.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
