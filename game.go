package main

import (
	"fmt"
)

func getNumPlayers() int {
	//Get number of players from user
	numPlayers := 0

	fmt.Println("How many players are in this game? Choose a number between 2 & 4 : ")
	fmt.Scan(&numPlayers)

	//check for valid choice
	for !IsValidNumPlayers(numPlayers) {
		fmt.Println("Invalid choice.")
		numPlayers = getNumPlayers()
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

func PlayGame() {

}
