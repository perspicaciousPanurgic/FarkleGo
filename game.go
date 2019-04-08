package main

import (
	"fmt"
)

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

func FindPoints(dice [6]int) int {
	// Initialize return value
	points := 0

	// Count number of times each value is rolled
	diceCount := CountDice(dice)

	// Check for 1's
	if diceCount[0] != 0 {
		num := diceCount[0]
		for num != 0 {
			if num > 2 {
				points += 1000
				num -= 3
				continue
			}
			points += num * 100
			num = 0
		}
	}

	// Check for 5's
	if diceCount[4] != 0 {
		num := diceCount[4]
		for num != 0 {
			if num > 2 {
				points += 500
				num -= 3
				continue
			}
			points += num * 50
			num = 0
		}
	}

	// Check for remaining 3 of a kind
	values := [...]int{1, 2, 3, 5}
	for _, value := range values {
		if diceCount[value] != 0 {
			num := diceCount[value]
			for num != 0 {
				if num > 2 {
					points += 100 * (value + 1)
					num -= 3
					continue
				}
				num = 0
			}
		}
	}

	return points
}

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

func NamePlayer(index int, name string, playerList []Player) {
	playerList[index].name = name
}

func PlayRound() {

}

func PlayGame() {

}
