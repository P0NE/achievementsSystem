package main

import "testing"

//Creation of a player test
var statisticTest = []statistic{
	statistic{
		AttackAttempt:    150,
		HitNumber:        80,
		TotalDamageDone:  450,
		Kill:             62,
		FirstHitKill:     12,
		Assist:           1450,
		TotalSpellCast:   20,
		TotalSpellDamage: 300,
		Time:             15,
	},
	statistic{
		AttackAttempt:    150,
		HitNumber:        120,
		TotalDamageDone:  530,
		Kill:             45,
		FirstHitKill:     14,
		Assist:           4050,
		TotalSpellCast:   250,
		TotalSpellDamage: 450,
		Time:             15,
	},
}

var gamesTest = []game{
	game{
		id:         0,
		NbPlayer:   8,
		Statistics: statisticTest[0],
		Team:       0,
	},
	game{
		id:         1,
		NbPlayer:   8,
		Statistics: statisticTest[1],
		Team:       1,
	},
}

var playerTest = player{
	GamerTag:     "TestMan",
	Games:        gamesTest,
	NbWin:        3045,
	NbLoss:       1055,
	Achievements: nil,
}

/*
 TU for BigBro achievement
 playerTest have more than 1000 assist so the achievement should be unlock
*/
func TestBigBro(t *testing.T) {
	isBigBro := bigBro(playerTest.Games)
	if !isBigBro {
		t.Errorf("BigBro achievement failed, expected to be unlock")
	}

}

/*
 TU for Veteran achievement
 playerTest have less than 1000 games so the achievement should not be unlock
*/
func TestVeteran(t *testing.T) {
	isVeteran := veteran(playerTest.Games)
	if isVeteran {
		t.Errorf("Veteran achievement failed, expected to not be unlock")
	}
}

/*
 TU for BigWinner achievement
 playerTest have more than 200 wins so the achievement should be unlock
*/
func TestBigWinner(t *testing.T) {
	isBigWinner := bigWinner(playerTest.NbWin)
	if !isBigWinner {
		t.Errorf("Big Winner achievement failed, expected to be unlock")
	}
}

/*
 TU for SharpShooter achievement
 player test have more 75% hits touched for its last game, so the achievement should be unlock
*/
func TestSharpShooter(t *testing.T) {
	isSharpShooter := sharpShooter(playerTest.getLastGame())
	if !isSharpShooter {
		t.Errorf("Sharp Shooter achievement failed, expected to be unlock")
	}
}

/*
 TU for Gandalf achievement
 player test have total spell damage which represent more than 50% of total damage, so the achievement gandalg should be unlock
*/
func TestGandalf(t *testing.T) {
	isGandalf := gandalf(playerTest.getLastGame())
	if !isGandalf {
		t.Errorf("Gandalf achievement failed, expected to be unlock")
	}
}

/*
 TU for Bruiser achievement
 player test have deal more than 500 damage in the last game, so the achievement Bruiser should be unlock
*/
func TestBruiser(t *testing.T) {
	isBruiser := bruiser(playerTest.getLastGame())
	if !isBruiser {
		t.Errorf("Bruiser achievement failed, expected to be unlock")
	}
}

/*
 TU for isAlreadyAchieved function
 PlayerTest avec unlock Bruiser and Gandalf but not BigBro
*/
func TestIsAlredayAchieved(t *testing.T) {
	var acheviementsTest = []string{"Bruiser", "Gandalf"}

	isAlreadyAchievedTest := isAlreadyAchieved(acheviementsTest, "Bruiser")
	if !isAlreadyAchievedTest {
		t.Errorf("Bruiser achievement already unlock, expected to be true")
	}

	isAlreadyAchievedTest = isAlreadyAchieved(acheviementsTest, "BigBro")
	if isAlreadyAchievedTest {
		t.Errorf("BigBro achievement not unlock, expected to be false")
	}
}
