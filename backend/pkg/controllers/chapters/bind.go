package chapters

type _gameFormat struct {
	Answers []string `json:"answers"`
	Matchs  []_match `json:"matchs"`
}

type _match struct {
	Word   string `json:"word"`
	Answer string `json:"answer"`
}
