package exporter

import (
	"net"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// ShipLogsFile sends accumulated logs to a remote server (shipper) using the configured protocol.
//
// Parameters:
//
//   - logs: a slice of maps where the outer key is the source and the inner map contains messages and their statistics (LogStats).
//
// Returns:
//
//   - error: nil if the transmission was successful or no address is configured; otherwise, returns connection, marshaling, or write errors.
//
// The function performs:
//
//   - If options.Config.ShipperConfig.Address is empty, it returns nil and does nothing.
//
//   - Establishes a connection using the protocol, address, and timeout specified in the configuration.
//
//   - Iterates through the slice of log maps:
//
//   - For each source (key) and its statistics (stat):
//
//     -Constructs a DataSentTcp structure with the origin and a data slice.
//
//     -For each message in that source, creates a TcpLogSent entry with count, firstSeen, lastSeen, and the message content.
//
//     -Serializes the DataSentTcp structure to JSON.
//
//     -Writes the JSON data to the connection.
//
//   - If any error occurs (connection, marshal, or write), it returns immediately.
//
//   - Upon completion, closes the connection and returns nil.
func ShipLogsFile(logs []map[string]map[string]*pipelines.LogStats) (err error) {
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
		data, err := ShipsIntegrations(messages)
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
