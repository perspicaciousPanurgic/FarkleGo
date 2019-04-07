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

	// Get the number of players
	numberOfPlayers := getNumPlayers()

	// FIXME: Import Game() func
	fmt.Println(numberOfPlayers)

}
