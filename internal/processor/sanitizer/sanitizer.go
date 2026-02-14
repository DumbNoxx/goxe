package sanitizer

import (
	"bytes"
)

// This function cleans the text by removing spaces and ignoring the words inside the filters
func Sanitizer(text []byte, idLog string) (newSlice []byte) {
	var infoWord []byte

	text = bytes.TrimSpace(text)
	text = reIpLogs.ReplaceAll(text, []byte(""))

	if len(idLog) > 0 {
		infoWord = SafeWord.ReplaceAll(text, []byte(""))
	} else {
		infoWord = text
	}

	textSanitize := bytes.ToLower(infoWord)
	infoText := reDates.ReplaceAll(textSanitize, []byte(""))
	cleanText := bytes.TrimSpace(reStatus.ReplaceAll(infoText, []byte("")))
	newSlice = ExtractLevelUpper(textSanitize)

	newSlice = append(newSlice, cleanText...)
	return newSlice
}
