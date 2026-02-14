package cluster

import (
	"bytes"

	"github.com/DumbNoxx/goxe/internal/processor/sanitizer"
)

func NormalizeLog(log []byte) []byte {
	return bytes.TrimSpace(sanitizer.Re.ReplaceAll(log, []byte("*")))
}
