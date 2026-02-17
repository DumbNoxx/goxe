package exporter

import (
	"encoding/json"
	"net"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

type DataSentTcp struct {
	Origin string       `json:"origin"`
	Data   []TcpLogSent `json:"tcpLogSent"`
}

type TcpLogSent struct {
	Count     int       `json:"count"`
	FirstSeen time.Time `json:"firstSeen"`
	LastSeen  time.Time `json:"lastSeen"`
	Message   string    `json:"message"`
}

// ShipLogs sends accumulated logs to a remote server (shipper) using the configured protocol.
//
// Parameters:
//
//   - logs: a map of maps where the outer key is the source and the inner map contains messages and their statistics.
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
//   - For each source in logs:
//
//     -Constructs a DataSentTcp structure with the origin and a slice of TcpLogSent.
//
//     -For each message in that source, creates a TcpLogSent with count, firstSeen, lastSeen, and the message content.
//
//     -Serializes the structure to JSON.
//
//     -Writes the data to the connection.
//
//   - If any error occurs (connection, marshal, or write), it returns immediately.
//
//   - Upon completion, closes the connection and returns nil.
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
	for key, messages := range logs {
		var DataSentTcps DataSentTcp
		DataSentTcps.Origin = key
		DataSentTcps.Data = make([]TcpLogSent, 0, len(messages))
		for msg, stats := range messages {
			var logEntry = TcpLogSent{
				Count:     stats.Count,
				FirstSeen: stats.FirstSeen,
				LastSeen:  stats.LastSeen,
				Message:   msg,
			}

			DataSentTcps.Data = append(DataSentTcps.Data, logEntry)
		}
		data, err := json.Marshal(DataSentTcps)
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
