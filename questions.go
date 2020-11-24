package english

import "strings"

var qPrefixes = []string{
	"what",
	"who",
	"where",
	"when",
	"why",
	"how",
}

func IsQuestion(s string) bool {
	s = strings.ToLower(s)
	for _, p := range qPrefixes {
		if strings.HasPrefix(s, p) &&
			strings.HasSuffix(s, "?") {
			return true
		}
	}
	return false
}
