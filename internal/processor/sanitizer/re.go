package sanitizer

import (
	"regexp"
	"strings"

	"github.com/DumbNoxx/Goxe/internal/processor/filters"
)

var (
	reStatus = regexp.MustCompile(filters.PatternsLogLevel)
	reDates  = regexp.MustCompile(strings.Join(filters.PatternsDate, "|"))
)
