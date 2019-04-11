//go main()
package main

import (
	"fmt"
	"io/ioutil"
)

// Global constants
const seconds int = 1

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

	// Play a game of Farkle
	PlayGame()
}
