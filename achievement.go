package main

const (
	// 75% hits touch
	SharpShooter = "SharpShooter"
	// 500 points of damage in one game
	Bruiser = "Bruiser"
	// more than 1000 games played
	Veteran = "Veteran"
	// more than 200 games win
	BigWinner = "BigWinner"
	// more than 1000 assist
	BigBro = "BigBro"
	// spell damage represent more than 50% of total damage in one game
	Gandalf = "Gandalf"
)

// If total spell damage represent more than 50% of total damage, achievement "Gandalf" unlock
func gandalf(p player) bool {
	if p.getLastGame().Statistics.TotalSpellDamage*100/p.getLastGame().Statistics.TotalDamageDone > 50 {
		return true
	}
	return false
}

//If total of assists is more than 5000
func bigBro(p player) bool {
	var totalAssist int
	for _, game := range p.Games {
		totalAssist += game.Statistics.Assist
	}
	if totalAssist >= 5000 {
		return true
	}
	return false

}

// If 75 % hits touched, achievement "SharpShooter" unlock
func sharpShooter(p player) bool {
	game := p.getLastGame()
	attackAttempt := game.Statistics.AttackAttempt
	hitNumber := game.Statistics.HitNumber
	pourcentageHit := (hitNumber * 100) / attackAttempt
	if pourcentageHit >= 75 {
		return true
	}
	return false
}

// If total damage done for the last game is more or else 500 pts, achievement "Bruiser" unlock
func bruiser(p player) bool {
	if p.getLastGame().Statistics.TotalDamageDone >= 500 {
		return true
	}
	return false
}

// If player played more than 1000 games, achievement "veteran" unlock
func veteran(p player) bool {
	if len(p.Games) >= 1000 {
		return true
	}
	return false
}

// If player have more than 200 wins, achievement "Big Winner" unlock
func bigWinner(p player) bool {
	if p.NbWin >= 200 {
		return true
	}
	return false
}

// Launch all achievements test
func launch(p player) player {
	if !isAlreadyAchieved(p.Achievements, SharpShooter) && sharpShooter(p) {
		p.Achievements = append(p.Achievements, SharpShooter)
	}

	if !isAlreadyAchieved(p.Achievements, Bruiser) && bruiser(p) {
		p.Achievements = append(p.Achievements, Bruiser)
	}

	if !isAlreadyAchieved(p.Achievements, Veteran) && veteran(p) {
		p.Achievements = append(p.Achievements, Veteran)
	}

	if !isAlreadyAchieved(p.Achievements, BigWinner) && bigWinner(p) {
		p.Achievements = append(p.Achievements, BigWinner)
	}

	if !isAlreadyAchieved(p.Achievements, BigBro) && bigBro(p) {
		p.Achievements = append(p.Achievements, BigBro)
	}

	if !isAlreadyAchieved(p.Achievements, Gandalf) && gandalf(p) {
		p.Achievements = append(p.Achievements, Gandalf)
	}
	return p
}

// Check if achievement already unlock
func isAlreadyAchieved(achievements []string, achievement string) bool {

	for _, a := range achievements {
		if a == achievement {
			return true
		}
	}
	return false
}
