package english

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

var irregularVerbs = map[string][]string{
	"see":   {"saw", "видеть"},
	"come":  {"came", "приходить"},
	"go":    {"went", "идти"},
	"know":  {"knew", "знать"},
	"give":  {"gave", "давать"},
	"take":  {"took", "брать"},
	"speak": {"spoke", "говорить"},
}
