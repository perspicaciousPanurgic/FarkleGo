package main

import (
	"fmt"
)

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

// Method to determine if a players turn is over
func EndRound(points int, numDice int, score int) bool {
	farkle := false

	// If no points were rolled then the round is over
	if points == 0 {
		farkle = true
		return farkle
	}

	// If < 1000 points rolled and score is < 1000  round is not over
	if points < 1000 && score < 1000 {
		return farkle
	}

	// If all six dice have scored points then player must roll again
	if numDice == 6 {
		return farkle
	}

	return farkle
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

func PlayRound() {

}

func PlayGame() {

}
