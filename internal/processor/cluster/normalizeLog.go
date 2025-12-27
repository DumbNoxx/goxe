package cluster

import (
	"regexp"
	"strings"
)

func NormalizeLog(log string) string {
	re := regexp.MustCompile(`\d+`)
	return strings.TrimSpace(re.ReplaceAllString(log, "*"))
}

