package cluster

import (
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/sanitizer"
)

func NormalizeLog(log string) string {
	return strings.TrimSpace(sanitizer.Re.ReplaceAllString(log, "*"))
}
