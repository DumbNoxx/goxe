package filters

import (
	"strings"

	"github.com/DumbNoxx/goxe/internal/options"
)

var Str *strings.Replacer

func LoadFiltersWord() {
	listIgnored := make([]string, 0, len(options.Config.PatternsWords)*2)
	for _, word := range options.Config.PatternsWords {
		listIgnored = append(listIgnored, word)
		listIgnored = append(listIgnored, "")
	}
	Str = strings.NewReplacer(listIgnored...)
}
