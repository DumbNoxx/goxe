package filters

import (
	"strings"
	"sync"

	"github.com/DumbNoxx/goxe/internal/options"
)

// Str is a global string replacer used to filter or remove specific
// words or patterns from logs. It is initialized using LoadFiltersWord.
var (
	currentReplacer *strings.Replacer
	mu sync.RWMutex
)

// LoadFiltersWord loads the words defined in the configuration (PatternsWords)
// and creates a strings.Replacer that replaces them with an empty string.
//
// The function performs:
//
//   - Iterates through options.Config.PatternsWords and constructs a list of pairs
//     (word to find, empty replacement) for strings.NewReplacer.
//   - Creates a new strings.Replacer with that list and assigns it to Str.
func LoadFiltersWord(getConfig options.ConfigProvider) {
	conf := getConfig()
	listIgnored := make([]string, 0, len(conf.PatternsWords)*2)
	for _, word := range conf.PatternsWords {
		listIgnored = append(listIgnored, word)
		listIgnored = append(listIgnored, "")
	}
	replacer := strings.NewReplacer(listIgnored...)
	mu.Lock()
	currentReplacer = replacer
	mu.Unlock()
}

func GetReplacer() *strings.Replacer {
	mu.RLock()
	defer mu.RUnlock()
	return currentReplacer 
}
