package exporter

import (
	"encoding/json"
	"log"

	"github.com/DumbNoxx/goxe/pkg/exporter"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// ShipsIntegrations transforms a map of log statistics into a JSON-encoded byte slice
// formatted for data transmission.
//
// Parameters:
//
//   - logs: map containing log statistics grouped by source and message.
//
// Returns:
//
//   - data: byte slice containing the JSON-marshaled data of the last processed source.
//   - err: error object if the JSON marshaling fails.
//
// The function performs:
//
//   - Iterates through the 'logs' map to process each source (key) and its associated messages.
//
//   - For each source:
//
//     -Initializes a 'DataSentTcp' structure and sets the 'Origin' field.
//
//     -Allocates a slice for 'TcpLogSent' entries with a capacity equal to the number of messages.
//
//     -Iterates through the messages to populate 'TcpLogSent' with 'Count', 'FirstSeen',
//     'LastSeen', and the message text.
//
//     -Appends each log entry to the 'DataSentTcps.Data' slice.
//
//     -Marshals the 'DataSentTcps' structure into a JSON byte slice.
//
//   - If an error occurs during marshaling, logs the error and returns an empty byte slice with the error.
//
//   - Returns the resulting byte slice and a nil error upon successful completion.
func ShipsIntegrations(logs map[string]map[string]*pipelines.LogStats) (data []byte, err error) {
	for key, messages := range logs {
		var DataSentTcps exporter.DataSent
		DataSentTcps.Origin = key
		DataSentTcps.Data = make([]exporter.LogSent, 0, len(messages))
		for msg, stats := range messages {
			var logEntry = exporter.LogSent{
				Count:     stats.Count,
				FirstSeen: stats.FirstSeen,
				LastSeen:  stats.LastSeen,
				Message:   msg,
			}
			DataSentTcps.Data = append(DataSentTcps.Data, logEntry)
		}
		data, err = json.Marshal(DataSentTcps)
		if err != nil {
			log.Println(err)
			return []byte{}, err
		}
	}
	return data, nil

}
