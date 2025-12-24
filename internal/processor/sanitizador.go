package processor

import (
	"regexp"
	"strings"
)

var reStatus = regexp.MustCompile(PatternsLogLevel)

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizador(text string) string {
	re := regexp.MustCompile(strings.Join(PatternsDate, "|"))
	textSanitize := strings.ToLower(strings.TrimSpace(text))
	infoText := re.ReplaceAllString(textSanitize, "")
	cleanText := strings.TrimSpace(reStatus.ReplaceAllString(infoText, ""))

	for _, word := range Ignored {
		if strings.Contains(word, cleanText) {
			return ""
		}
	}

	return extractLevelUpper(textSanitize) + cleanText
}

func extractLevelUpper(log string) string {
	status := reStatus.FindString(log)

	if status == "" {
		return ""
	}
	return strings.ToUpper(status)
}
