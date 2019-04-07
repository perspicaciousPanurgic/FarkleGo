package main

import (
	"fmt"
	"log"
)

func getNumPlayers() int {
	//Get number of players from user
	numPlayers := 0

	fmt.Println("How many players are in this game? Choose a number between 2 & 4 : ")
	_, err := fmt.Scanf("&d", &numPlayers)
	if err != nil {
		log.Fatal(err)
	}

	//check for valid choice
	for numPlayers < 2 && numPlayers > 4 {
		fmt.Println("Invalid choice. You must choose a number between 2 & 4 : ")
		_, err := fmt.Scanf("&d", &numPlayers)
		if err != nil {
			log.Fatal(err)
		}
	}

	return numPlayers
}

func isValidNumPlayers() {

}
