package exporter

import (
	"encoding/json"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// Shipper defines the contract for transforming aggregated log statistics
// into a format suitable for external transmission.
//
// Any implementation of this interface is responsible for taking the internal
// log representation and encoding it according to the requirements of the
// specific output destination (e.g., JSON, CloudWatch events, XML).
type Shipper interface {
	// PrepareShip transforms a nested map of LogStats into a byte slice.
	//
	// Parameters:
	//   - logs: A map where the first key is the log source and the second
	//           key is the sanitized log message.
	//
	// Returns:
	//   - data: The formatted byte slice ready to be sent over the wire.
	//   - err:  An error if the transformation or encoding process fails.
	PrepareShip(logs map[string]map[string]*pipelines.LogStats) (data []byte, err error)
}

// JsonManager implements the exporter.Shipper interface to handle log
// transformation into JSON format.
//
// This struct acts as a concrete provider for generic log exports,
// focusing on converting internal LogStats maps into valid JSON byte slices
// for network transmission or integration calls.
//
// It is used by the system when the "generic" destination is configured.
type JsonManager struct{}

// PrepareShip transforms a map of log statistics into a JSON-encoded byte slice
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
//     -Initializes a 'DataSent' structure and sets the 'Origin' field.
//
//     -Allocates a slice for 'LogSent' entries with a capacity equal to the number of messages.
//
//     -Iterates through the messages to populate 'LogSent' with 'Count', 'FirstSeen',
//     'LastSeen', and the message text.
//
//     -Appends each log entry to the 'DataSent.Data' slice.
//
//     -Marshals the 'DataSent' structure into a JSON byte slice.
//
//   - If an error occurs during marshaling, returns an empty byte slice with the error.
//
//   - Returns the resulting byte slice and a nil error upon successful completion.
func (shipper *JsonManager) PrepareShip(logs map[string]map[string]*pipelines.LogStats) (data []byte, err error) {
	for key, messages := range logs {
		var DataSent DataSent
		DataSent.Origin = key
		DataSent.Data = make([]LogSent, 0, len(messages))
		for msg, stats := range messages {
			var logEntry = LogSent{
				Count:     stats.Count,
				FirstSeen: stats.FirstSeen,
				LastSeen:  stats.LastSeen,
				Message:   msg,
			}
			DataSent.Data = append(DataSent.Data, logEntry)
		}
		data, err = json.Marshal(DataSent)
		if err != nil {
			return []byte{}, err
		}
	}
	return data, nil
}
