package english

import "strings"

var questions = map[string][]string{
	"what":  {"что", "какой"},
	"who":   {"кто"},
	"where": {"где", "куда"},
	"when":  {"когда"},
	"why":   {"почему", "зачем"},
	"how":   {"как"},
	"do":    {"делать"},
	"does":  {"делает"},
	"did":   {"делал"},
	"will":  {"буду"},
	"am":    {"есть"},
	"is":    {"есть"},
	"are":   {"есть"},
	"was":   {"был"},
	"were":  {"были"},
}

func IsQuestion(s string) bool {
	s = strings.ToLower(s)
	for q := range questions {
		if strings.HasPrefix(s, q) &&
			strings.HasSuffix(s, "?") {
			return true
		}
	}
	return false
}
