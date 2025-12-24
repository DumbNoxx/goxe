package processor

import "strings"

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizador(text string) string {
	cleanText := strings.ToLower(strings.TrimSpace(text))

	for _, word := range Ignored {
		if strings.Contains(word, cleanText) {
			return ""
		}
	}

	return cleanText
}
