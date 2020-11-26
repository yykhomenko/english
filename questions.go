package english

import "strings"

var questions = map[string][]string{
	"what":  {"что", "какой"},
	"who":   {"кто"},
	"where": {"где", "куда"},
	"when":  {"когда"},
	"why":   {"почему", "зачем"},
	"how":   {"как"},
}

func IsQuestion(s string) bool {
	s = strings.ToLower(s)

	for q := range questions {
		if strings.HasPrefix(s, q) &&
			strings.HasSuffix(s, "?") {
			return true
		}
	}

	for q := range do {
		if strings.HasPrefix(s, q) &&
			strings.HasSuffix(s, "?") {
			return true
		}
	}

	for q := range toBe {
		if strings.HasPrefix(s, q) &&
			strings.HasSuffix(s, "?") {
			return true
		}
	}

	return false
}
