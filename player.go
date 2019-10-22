package main

// Structure representation of a player
type player struct {
	GamerTag     string
	Games        []game
	NbWin        int
	NbLoss       int
	Achievements []string
}

// Function which send the last game played
func (p *player) getLastGame() game {
	return p.Games[len(p.Games)-1]
}
