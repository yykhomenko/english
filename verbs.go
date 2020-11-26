package english

// правильные глаголы
var verbsRegular = map[string][]string{
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
var verbsIrregular = map[string][]string{
	"see":   {"saw", "видеть"},
	"come":  {"came", "приходить"},
	"go":    {"went", "идти"},
	"know":  {"knew", "знать"},
	"give":  {"gave", "давать"},
	"take":  {"took", "брать"},
	"speak": {"spoke", "говорить"},
}

// глагол "делать"
var verbsDo = map[string][]string{
	"do":   {"i", "you", "we", "they"},
	"does": {"he", "she", "it"},
	"did":  {"i", "you", "we", "they", "he", "she", "it"},
}

// глагол "быть"
var verbsToBe = map[string][]string{
	"will": {"i", "you", "we", "they", "he", "she", "it"},
	"am":   {"i"},
	"is":   {"he", "she", "it"},
	"are":  {"you", "we", "they"},
	"was":  {"i", "he", "she", "it"},
	"were": {"you", "we", "they"},
}
