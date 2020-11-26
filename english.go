package english

func NumOfSentences() int {
	return len(adverbs) *
		len(prepositions) *
		len(pronounsWho) *
		len(pronounsWhom) *
		len(pronounsWhose) *
		len(questions) *
		len(verbsRegular) *
		len(verbsIrregular) *
		len(verbsDo) *
		len(verbsToBe)
}
