package sanitizer

import "strings"

func extractLevelUpper(log string) string {
	status := reStatus.FindString(log)

	if status == "" {
		return ""
	}
	return strings.ToUpper(status)
}
