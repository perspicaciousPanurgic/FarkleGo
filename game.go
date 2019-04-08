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
	points := 0

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
