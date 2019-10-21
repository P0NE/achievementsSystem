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
func gandalf(lastGame game) bool {
	if lastGame.Statistics.TotalSpellDamage*100/lastGame.Statistics.TotalDamageDone > 50 {
		return true
	}
	return false
}

//If total of assists is more than 5000
func bigBro(games []game) bool {
	var totalAssist int
	for _, game := range games {
		totalAssist += game.Statistics.Assist
	}
	if totalAssist >= 5000 {
		return true
	}
	return false

}

// If 75 % hits touched, achievement "SharpShooter" unlock
func sharpShooter(lastGame game) bool {
	game := lastGame
	attackAttempt := game.Statistics.AttackAttempt
	hitNumber := game.Statistics.HitNumber
	pourcentageHit := (hitNumber * 100) / attackAttempt
	if pourcentageHit >= 75 {
		return true
	}
	return false
}

// If total damage done for the last game is more or else 500 pts, achievement "Bruiser" unlock
func bruiser(lastGame game) bool {
	if lastGame.Statistics.TotalDamageDone >= 500 {
		return true
	}
	return false
}

// If player played more than 1000 games, achievement "veteran" unlock
func veteran(games []game) bool {
	if len(games) >= 1000 {
		return true
	}
	return false
}

// If player have more than 200 wins, achievement "Big Winner" unlock
func bigWinner(nbWin int) bool {
	if nbWin >= 200 {
		return true
	}
	return false
}

/*
Launch all achievements test
Todo: use other method for launching check achievements
*/
func launch(p player) player {
	if !isAlreadyAchieved(p.Achievements, SharpShooter) && sharpShooter(p.getLastGame()) {
		p.Achievements = append(p.Achievements, SharpShooter)
	}

	if !isAlreadyAchieved(p.Achievements, Bruiser) && bruiser(p.getLastGame()) {
		p.Achievements = append(p.Achievements, Bruiser)
	}

	if !isAlreadyAchieved(p.Achievements, Veteran) && veteran(p.Games) {
		p.Achievements = append(p.Achievements, Veteran)
	}

	if !isAlreadyAchieved(p.Achievements, BigWinner) && bigWinner(p.NbWin) {
		p.Achievements = append(p.Achievements, BigWinner)
	}

	if !isAlreadyAchieved(p.Achievements, BigBro) && bigBro(p.Games) {
		p.Achievements = append(p.Achievements, BigBro)
	}

	if !isAlreadyAchieved(p.Achievements, Gandalf) && gandalf(p.getLastGame()) {
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
