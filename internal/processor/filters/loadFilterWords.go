package filters

import (
	"strings"

	"github.com/DumbNoxx/goxe/internal/options"
)

// Str is a global string replacer used to filter or remove specific
// words or patterns from logs. It is initialized using LoadFiltersWord.
var Str *strings.Replacer

// LoadFiltersWord loads the words defined in the configuration (PatternsWords)
// and creates a strings.Replacer that replaces them with an empty string.
//
// The function performs:
//
//   - Iterates through options.Config.Pa and constructs a list of pairs
//     (word to find, empty replacement) for strings.NewReplacer.
//   - Creates a new strings.Replacer with that list and assigns it to Str.
func LoadFiltersWord() {
	listIgnored := make([]string, 0, len(options.Config.PatternsWords)*2)
	for _, word := range options.Config.PatternsWords {
		listIgnored = append(listIgnored, word)
		listIgnored = append(listIgnored, "")
	}
	Str = strings.NewReplacer(listIgnored...)
}
