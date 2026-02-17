package cluster

import (
	"bytes"

	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
)

// NormalizeLog normalizes a log by replacing specific patterns and trimming spaces.
//
// Parameters:
//   - log: []byte containing the log content to be normalized.
//
// Returns:
//   - []byte: the modified log, with patterns replaced by '*' and leading/trailing spaces removed
//
// The function performs:
//   - Applies a regular expression (sanitizer.Re) to replace all matches with '*'.
//   - Then removes leading and trailing whitespace using bytes.TrimSpace.
func NormalizeLog(log []byte) []byte {
	return bytes.TrimSpace(sanitizer.Re.ReplaceAll(log, []byte("*")))
}
