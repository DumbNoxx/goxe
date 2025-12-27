package sanitizer

import (
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/filters"
)

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizer(text string) string {
	textSanitize := strings.ToLower(strings.TrimSpace(text))
	infoText := reDates.ReplaceAllString(textSanitize, "")
	cleanText := strings.TrimSpace(reStatus.ReplaceAllString(infoText, ""))

	for _, word := range filters.Ignored {
		if strings.Contains(word, cleanText) {
			return ""
		}
	}

	return extractLevelUpper(textSanitize) + cleanText
}
