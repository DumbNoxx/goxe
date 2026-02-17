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

// SafeWordFunc constructs a regular expression that searches for the pattern "<idLog>_<numeric ID>".
//
// Parameters:
//
//   - word: byte slice containing the base identifier (e.g., the value of options.Config.IdLog).
//
// Returns:
//
//   - *regexp.Regexp: a compiled regular expression that recognizes the identifier followed by an underscore
//     and a sequence of digits (based on filters.PatternsIdLogs).
//
// The function performs:
//
//   - Builds a string by concatenating the identifier (word converted to string) with an underscore.
//   - Escapes any special characters using regexp.QuoteMeta to ensure they are interpreted literally.
//   - Appends the digit pattern from filters.PatternsIdLogs.
//   - Compiles and returns the resulting regular expression.
func SafeWordFunc(word []byte) *regexp.Regexp {
	var newWord strings.Builder
	fmt.Fprint(&newWord, string(word))
	fmt.Fprint(&newWord, "_")
	return regexp.MustCompile(regexp.QuoteMeta(newWord.String()) + filters.PatternsIdLogs)
}
