package main

/*
Function launching the achievements system
This function can be launched by goroutines for each games.
Use mutex for checking that achievements are not launch for the same player in the same time
*/
func runAchievements(players []player) {

	for index, player := range players {
		players[index] = launch(player)
	}
}
