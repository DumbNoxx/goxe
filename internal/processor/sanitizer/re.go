package sanitizer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor/filters"
)

var (
	reStatus = regexp.MustCompile(filters.PatternsLogLevel)
	reDates  = regexp.MustCompile(strings.Join(filters.PatternsDate, "|"))
	reIpLogs = regexp.MustCompile(filters.PatternIpLogs)
	Re       = regexp.MustCompile(`\d+`)
	SafeWord = SafeWordFunc([]byte(options.Config.IdLog))
)

func SafeWordFunc(word []byte) *regexp.Regexp {
	var newWord strings.Builder
	fmt.Fprint(&newWord, string(word))
	fmt.Fprint(&newWord, "_")
	return regexp.MustCompile(regexp.QuoteMeta(newWord.String()) + filters.PatternsIdLogs)
}
