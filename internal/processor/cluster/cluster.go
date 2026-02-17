package cluster

import "github.com/DumbNoxx/goxe/internal/processor/sanitizer"

// Cluster processes a log by applying sanitization and normalization.
//
// Parameters:
//   - log: original log content in bytes.
//   - idLog: unique identifier for the log.
//
// Returns:
//   - []byte: the resulting log after sanitizing and normalization.
//
// The function performs:
//   - Calls sanitizer.Sanitizer with log and idLog to clean the text.
//   - Then calls NormalizeLog on the result to unity the format.
func Cluster(log []byte, idLog string) []byte {
	text := sanitizer.Sanitizer(log, idLog)
	normalizeLog := NormalizeLog(text)

	return normalizeLog
}
