package main

import (
	"fmt"
)

//Global constants
const MaxDice = 6
const PointsToWin = 2000

// Method to solicit Y or N from user
func GetBinaryChoice() string {
	userInput := "Y"

	fmt.Println("Enter either Y or N : ")
	fmt.Scan(&userInput)

	//check for valid choice
	for !isValidBinaryChoice(userInput) {
		fmt.Println("Invalid choice.")
		userInput = GetBinaryChoice()
	}

	return userInput
}

// Check if user input is valid Y or N
func isValidBinaryChoice(input string) bool {
	valid := false
	if input == "Y" || input == "N" {
		valid = true
	}
	return valid
}

// Convert user input of Y or N to boolean
func ConvertBinaryChoice(userInput string) bool {
	choice := false

	if userInput == "Y" {
		choice = true
	}

	return choice
}

// Method to track whose turn it is
func NextActivePlayer(currentPlayer int, numPlayers int) int {
	// Declare return variable equal to current player
	next := currentPlayer

	// Increment to next player
	next += 1

	// If player is the last player return to first player
	if next > (numPlayers - 1) {
		next = 0
	}

	return next
}

// Method to execute the effects of a player choosing to hold point scoring dice
func KeepPoints(points int, player *Player) {
	player.AddScore(points)
}

// Method to evaluate dice rolled for points. Returns points kept by player and numDice used
// in a rollResult struct
func FindPoints(dice [6]int) rollResult {

	// Initialize return values
	points := 0
	numDice := 0

	// Count number of times each value is rolled
	diceCount := CountDice(dice)

	// Check for 1's
	if diceCount[0] != 0 {
		pts := 0
		d := 0
		num := diceCount[0]
		for num != 0 {
			if num > 2 {
				pts += 1000
				num -= 3
				d += 3
				continue
			}
			pts += num * 100
			d += num
			num = 0
		}
		// Ask user if they want to keep points
		fmt.Printf("You rolled %d 1's worth %d points. Do you want to hold them? \n", diceCount[0], pts)
		choice := ConvertBinaryChoice(GetBinaryChoice())

		// add to active player score
		if choice {
			points += pts
			numDice += d
		}
	}

	// Check for 5's
	if diceCount[4] != 0 {
		pts := 0
		d := 0
		num := diceCount[4]
		for num != 0 {
			if num > 2 {
				pts += 500
				num -= 3
				d += 3
				continue
			}
			pts += num * 50
			d += num
			num = 0
		}
		// Ask user if they want to keep points
		fmt.Printf("You rolled %d 5's worth %d points. Do you want to hold them? \n", diceCount[4], pts)
		choice := ConvertBinaryChoice(GetBinaryChoice())

		// add to active player score
		if choice {
			points += pts
			numDice += d
		}
	}

	// Check for remaining 3 of a kind
	values := [...]int{1, 2, 3, 5}
	for _, value := range values {
		if diceCount[value] > 2 {
			pts := 0
			d := 0
			num := diceCount[value]
			for num != 0 {
				if num > 2 {
					pts += 100 * (value + 1)
					num -= 3
					d += 3
					continue
				}
				num = 0
			}
			// Ask user if they want to keep points
			fmt.Printf("You rolled 3 %d's worth %d points. Do you want to hold them? \n", value+1, pts)
			choice := ConvertBinaryChoice(GetBinaryChoice())

			// add to active player score
			if choice {
				points += pts
				numDice += d
			}
		}
	}
	result := rollResult{points, numDice}
	return result
}

// Method to solicit user input on how many players there will be
func GetNumPlayers() int {
	//Get number of players from user
	numPlayers := 0

	fmt.Println("How many players are in this game? Choose a number between 2 & 4 : ")
	fmt.Scan(&numPlayers)

	//check for valid choice
	for !IsValidNumPlayers(numPlayers) {
		fmt.Println("Invalid choice.")
		numPlayers = GetNumPlayers()
	}

	return numPlayers
}

// Method to validate number of players is between 2 and 4
func IsValidNumPlayers(num int) bool {
	isValid := true

	if num < 2 {
		isValid = false
	}

	if num > 4 {
		isValid = false
	}

	return isValid
}

// Method to name each active player
func NamePlayer(index int, name string, playerList []Player) {
	playerList[index].name = name
}

func SetupGame(numPlayers int) []Player {
	// Initialize array of players
	playerList := []Player{
		Player{1, "Adam", 0},
		Player{2, "Bob", 0},
		Player{3, "Chris", 0},
		Player{4, "Dave", 0},
	}

	// name the players
	for i := 0; i < numPlayers; i++ {
		//Get name of player from user
		name := ""
		fmt.Printf("What is the name of Player %d : \n", i+1)
		fmt.Scan(&name)

		// Change player name
		NamePlayer(i, name, playerList)
	}

	return playerList
}

func PlayGame() {

	// Setup the Game
	numPlayers := GetNumPlayers()
	playerList := SetupGame(numPlayers)
	activePlayer := -1

	// Initialize Dice Array & eligible dice to roll
	activeDice := [6]int{}
	numDice := MaxDice

	// State who is the active player
	activePlayer = NextActivePlayer(activePlayer, numPlayers)
	fmt.Printf("\nThe active player is Player %d : %v\n", activePlayer+1, playerList[activePlayer].name)

	// Round loop
	farkle := false
	for farkle == false {
		// Roll Dice
		activeDice = Roll(activeDice, numDice)

		// Print Dice
		PrintActiveDice(activeDice)

		// Find points
		result := FindPoints(activeDice)

		// Print results
		//fmt.Printf("%v now has %d points.", playerList[activePlayer].name, playerList[activePlayer].score)
		//fmt.Printf("%d dice may be rolled again.", MaxDice-result.numDice)

		// End of round evaluation to determine if round is over
		farkle = EndRound(result.points, result.numDice, playerList[activePlayer].score, &playerList[activePlayer])

		// Determine dice to roll
		if result.numDice != 6 {
			numDice = numDice - result.numDice
		}
	}
}

// Method to determine if a players turn is over
func EndRound(points int, numDice int, score int, player *Player) bool {
	farkle := true

	// If no points were rolled then the round is over
	if points == 0 {
		return farkle
	}

	// If all six dice have scored points then player must roll again
	if numDice == 6 {
		farkle = false
		return farkle
	}

	// If player score is less than 1000 and they have not rolled 1000 points they must roll again
	if score < 1000 {
		if points < 1000 {
			farkle = false
			return farkle
		}
	}

	// Ask if player wants to keep points and end turn
	fmt.Printf("Do you want to keep %d points or roll %d dice?\n", points, MaxDice-numDice)
	choice := ConvertBinaryChoice(GetBinaryChoice())

	// If player chooses to roll
	if choice == false {
		farkle = false
	}

	// Add points to score
	player.AddScore(points)

	return farkle
}

func PlayRound() {

}
