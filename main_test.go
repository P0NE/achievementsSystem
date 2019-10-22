package main

import (
	"fmt"
	"math/rand"
	"testing"
)

var players = make([]player, 8)

//Driver functions for emulate game
func createGame(id int) game {
	//for each games create random statistics
	return game{
		id:         id,
		NbPlayer:   8,
		Statistics: createStatistic(),
		Team:       rand.Intn(2),
	}

}

//Driver functions for emulate history
func createHistory(numberGame int) []game {
	//create 5 game for each players
	games := make([]game, 0, numberGame)
	for i := 0; i < numberGame; i++ {
		games = append(games, createGame(i))
	}

	return games

}

//Driver functions for emulate statistic
func createStatistic() statistic {
	//create random statistic for a game of a player

	return statistic{
		AttackAttempt:    rand.Intn(100),
		HitNumber:        rand.Intn(90),
		TotalDamageDone:  rand.Intn(800),
		Kill:             rand.Intn(50),
		FirstHitKill:     rand.Intn(15),
		Assist:           rand.Intn(50),
		TotalSpellCast:   rand.Intn(100),
		TotalSpellDamage: rand.Intn(800),
		Time:             rand.Intn(30),
	}

}

//launch driver
func launchDriver(numberGames int) {
	//create 8 players
	players[0] = player{
		GamerTag: "faker",
		NbWin:    0,
		NbLoss:   0,
	}
	players[1] = player{
		GamerTag: "Uzi",
		NbWin:    0,
		NbLoss:   0,
	}
	players[2] = player{
		GamerTag: "Soaz",
		NbWin:    0,
		NbLoss:   0,
	}
	players[3] = player{
		GamerTag: "Peke",
		NbWin:    0,
		NbLoss:   0,
	}
	players[4] = player{
		GamerTag: "Ocelote",
		NbWin:    0,
		NbLoss:   0,
	}
	players[5] = player{
		GamerTag: "CoreJJ",
		NbWin:    0,
		NbLoss:   0,
	}
	players[6] = player{
		GamerTag: "Roy",
		NbWin:    0,
		NbLoss:   0,
	}
	players[7] = player{
		GamerTag: "Walshy",
		NbWin:    0,
		NbLoss:   0,
	}

	// for each players, create history
	for index := range players {
		players[index].Games = createHistory(numberGames)
		players[index].NbWin = rand.Intn(numberGames)
		players[index].NbLoss = numberGames - players[index].NbWin
	}

}

// Launch the driver and after the achievements system
func TestAchievements(t *testing.T) {

	launchDriver(rand.Intn(1500))
	runAchievements(players)

	for _, player := range players {
		fmt.Println(player.Achievements)
	}

}
