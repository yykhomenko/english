package english

var regularVerbs = map[string][]string{
	"love":   {"любить"},
	"live":   {"жить"},
	"work":   {"работать"},
	"open":   {"открывать"},
	"close":  {"закрывать"},
	"start":  {"начинать"},
	"finish": {"заканчивать"},
}

var irregularVerbs = map[string][]string{
	"see":  {"saw", "видеть"},
	"come": {"came", "приходить"},
	"go":   {"went", "идти"},
	"know": {"knew", "знать"},
}
