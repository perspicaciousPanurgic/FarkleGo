package main

import (
	"fmt"
	"io/ioutil"
)

func Greetings() string {
	return "Let's Play Farkle!"
}

func PrintRules() {
	file, err := ioutil.ReadFile("farkleRules.txt")
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

	// Play the game
	// FIXME: Import Game() func

}
