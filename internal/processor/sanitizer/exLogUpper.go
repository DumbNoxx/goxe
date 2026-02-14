package sanitizer

import (
	"bytes"
)

func ExtractLevelUpper(log []byte) []byte {
	status := reStatus.Find(log)

	if len(status) <= 0 {
		return nil
	}
	return bytes.ToUpper(status)
}
