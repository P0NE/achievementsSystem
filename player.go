package main

type player struct {
	GamerTag     string
	Games        []game
	NbWin        int
	NbLoss       int
	Achievements []string
}

func (p *player) getLastGame() game {
	return p.Games[len(p.Games)-1]
}
