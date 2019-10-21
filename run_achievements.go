package main

func runAchievements(players []player) {

	for index, player := range players {
		players[index] = launch(player)
	}
}
