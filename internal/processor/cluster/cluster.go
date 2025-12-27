package cluster

import "github.com/DumbNoxx/Goxe/internal/processor/sanitizer"

func Cluster(log string) string {
	text := sanitizer.Sanitizer(log)
	normalizeLog := NormalizeLog(text)

	return normalizeLog
}
