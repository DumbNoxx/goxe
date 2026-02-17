package sanitizer

import (
	"bytes"
)

// ExtractLevelUpper extracts the log level (e.g.,'ERROR', 'info') from a log parameter
//
// Parameters:
//
// - log: []byte containing the log content.
//
// Returns:
//
//   - []byte: the found level in uppercase (e.g.,[]byte("ERROR")).
//     If no level is found it returns nil.
//
// The function performs:
//
//   - Applies reStatus.Find(log) to search the first match of the pattern.
//   - If the math is empty (len <= 0), it returns nil.
//   - If found, it converts the result to uppercase using bytes.Upper and returns it.
func ExtractLevelUpper(log []byte) []byte {
	status := reStatus.Find(log)

	if len(status) <= 0 {
		return nil
	}
	return bytes.ToUpper(status)
}
