package cluster

import "github.com/DumbNoxx/goxe/internal/processor/sanitizer"

func Cluster(log []byte, idLog string) []byte {
	text := sanitizer.Sanitizer(log, idLog)
	normalizeLog := NormalizeLog(text)

	return normalizeLog
}
