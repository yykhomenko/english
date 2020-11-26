package english

// правильные глаголы
var regularVerbs = map[string][]string{
	"love":   {"любить"},
	"live":   {"жить"},
	"work":   {"работать"},
	"open":   {"открывать"},
	"close":  {"закрывать"},
	"start":  {"начинать"},
	"finish": {"заканчивать"},
	"ask":    {"просить", "спрашивать"},
	"answer": {"отвечать"},
	"help":   {"помогать"},
	"like":   {"нравится"},
	"want":   {"хотеть"},
}

// неправильные глаголы
var irregularVerbs = map[string][]string{
	"see":   {"saw", "видеть"},
	"come":  {"came", "приходить"},
	"go":    {"went", "идти"},
	"know":  {"knew", "знать"},
	"give":  {"gave", "давать"},
	"take":  {"took", "брать"},
	"speak": {"spoke", "говорить"},
}

var do = map[string][]string{
	"will": {"i", "you", "we", "they", "he", "she", "it"},
	"do":   {"i", "you", "we", "they"},
	"does": {"he", "she", "it"},
	"did":  {"i", "you", "we", "they", "he", "she", "it"},
}

var toBe = map[string][]string{
	"will": {"i", "you", "we", "they", "he", "she", "it"},
	"am":   {"i"},
	"is":   {"he", "she", "it"},
	"are":  {"you", "we", "they"},
	"was":  {"i", "he", "she", "it"},
	"were": {"you", "we", "they"},
}
