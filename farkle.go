//go main()
package main

import (
	"fmt"
	"io/ioutil"
)

//Global constants
const MaxDice = 6
const PointsToWin = 2000

func Greetings() string {
	return "Let's Play Farkle!"
}

func PrintRules() {
	file, err := ioutil.ReadFile("FarkleRules.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(file) // convert file to string

	fmt.Println(str) // print string
}

func main() {
	// Greet the user
	println(Greetings())

	// Print Farkle rules
	PrintRules()

	// Initialize array of default players
	playerList := []Player{
		Player{1, "Adam", 0},
		Player{2, "Bob", 0},
		Player{3, "Chris", 0},
		Player{4, "Dave", 0},
	}

	// Initialize global game variables
	activePlayer := -1

	// Get the number of players
	numPlayers := GetNumPlayers()

	// name the players
	for i := 0; i < numPlayers; i++ {
		//Get name of player from user
		name := ""
		fmt.Printf("What is the name of Player %d : \n", i+1)
		fmt.Scan(&name)

		// Change player name
		NamePlayer(i, name, playerList)
	}

	// Initialize Dice Array & eligible dice to roll
	activeDice := [6]int{}
	numDice := MaxDice
	//farkle := false

	// State who is the active player
	activePlayer = NextActivePlayer(activePlayer, numPlayers)
	fmt.Printf("\nThe active player is Player %d : %v\n", activePlayer+1, playerList[activePlayer].name)

	// Roll Dice
	activeDice = Roll(activeDice, numDice)

	// Print Dice
	PrintActiveDice(activeDice)

	// Find points
	result := FindPoints(activeDice)

	// Add to active player's score
	playerList[activePlayer].AddScore(result.points)

	// Print results
	fmt.Printf("%v now has %d points.", playerList[activePlayer].name, playerList[activePlayer].score)
	fmt.Printf("%d dice may be rolled again.", MaxDice-result.numDice)
}
