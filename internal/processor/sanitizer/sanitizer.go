package sanitizer

import (
	"bytes"
)

// Sanitizer cleans a log by removing spaces, IP addresses, dates, log levels, and,
// optionally, specific identifiers, while extracting the log level in uppercase.
//
// Parameters:
//
//   - text: byte slice containing the original log content.
//   - idLog: log identifier (if not empty, it is used to remove patterns
//     of the type "<idLog>_<numeric ID>" via the SafeWord regular expression).
//
// Returns:
//
//   - []byte: the processed log, consisting of the extracted log level (in uppercase)
//     followed by the clean text (without extreme spaces, IPs, dates, or levels).
//
// The function performs:
//
//   - Trims leading and trailing spaces from 'text' using bytes.TrimSpace.
//   - Replaces all IP addresses (reIpLogs) with an empty string.
//   - If idLog is not empty, it applies SafeWord.ReplaceAll to remove
//     identifier patterns; if empty, it keeps 'text' as is.
//   - Converts the result to lowercase (bytes.ToLower).
//   - Removes dates (reDates) and log levels (reStatus) by replacing them with empty strings.
//   - Trims spaces again with bytes.TrimSpace.
//   - Extracts the log level in uppercase by calling ExtractLevelUpper on the
//     lowercase version of the text (before levels are removed).
//   - Concatenates the extracted level with the clean text and returns the result.
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
