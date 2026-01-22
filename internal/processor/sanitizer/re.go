package sanitizer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/filters"
)

var (
	reStatus = regexp.MustCompile(filters.PatternsLogLevel)
	reDates  = regexp.MustCompile(strings.Join(filters.PatternsDate, "|"))
	reIpLogs = regexp.MustCompile(filters.PatternIpLogs)
)

func SafeWord(word string) *regexp.Regexp {
	var newWord strings.Builder
	fmt.Fprint(&newWord, word)
	fmt.Fprint(&newWord, "_")
	return regexp.MustCompile(regexp.QuoteMeta(newWord.String()) + filters.PatternsIdLogs)
}
