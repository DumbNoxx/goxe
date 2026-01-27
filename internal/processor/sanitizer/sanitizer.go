package sanitizer

import (
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/filters"
)

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizer(text string, idLog string) string {
	var infoWord string

	text = strings.TrimSpace(text)
	text = reIpLogs.ReplaceAllString(text, "")

	if len(idLog) > 0 {
		infoWord = SafeWord.ReplaceAllString(text, "")
	} else {
		infoWord = text
	}

	textSanitize := strings.ToLower(infoWord)
	infoText := reDates.ReplaceAllString(textSanitize, "")
	cleanText := strings.TrimSpace(reStatus.ReplaceAllString(infoText, ""))

	for _, word := range filters.Ignored {
		if strings.Contains(word, cleanText) {
			return ""
		}
	}

	return ExtractLevelUpper(textSanitize) + cleanText
}
