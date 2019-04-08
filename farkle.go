//go main()
package main

import (
	"fmt"
	"io/ioutil"
)

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

	// State who is the active player
	activePlayer = NextActivePlayer(activePlayer, numPlayers)
	fmt.Printf("\nThe active player is Player %d : %v\n", activePlayer+1, playerList[activePlayer].name)

	// Initialize Dice Arrays
	activeDice := [6]int{}
	heldDice := [6]int{}

	// Roll Dice
	activeDice = Roll(activeDice, 6)

	// Print Dice
	PrintActiveDice(activeDice)
	PrintHeldDice(heldDice)

	// Find points
	points := FindPoints(activeDice)
	fmt.Printf("The dice roll is worth %d points.", points)
}
