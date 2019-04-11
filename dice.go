package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct to hold roll information for FindPoints()
type rollResult struct {
	points   int
	keptDice int
}

//Roll to generate random numbers between 1-6
func Roll(dice *[6]int, numDice *int) *[6]int {

	// generate random source seed using time
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// fill array with zeroes to ensure clean start
	for i := range dice {
		dice[i] = 0
	}

	// start at index 0 and fill array with random values between 1 & 6 until numDice index
	// this simulates rolling numDice number of dice
	for i := 0; i < *numDice; i++ {
		dice[i] = r1.Intn(5) + 1
	}

	return dice
}

// Returns [6]int array which contains the number of times each value occurs given a dice roll
func CountDice(dice *[6]int) [6]int {
	// Initialize new array to store dice count
	count := [6]int{}

	// Iterate through dice roll and count number of time each value occurs
	for i := 0; i < 6; i++ {
		num := 0
		for j := 0; j < 6; j++ {
			if dice[j] == (i + 1) {
				num++
			}
		}
		count[i] = num
	}

	return count
}

// Method to print the dice roll
func PrintActiveDice(dice *[6]int) {
	// Wait for 1 second
	slowDown(seconds)
	// Print dice array
	fmt.Print("\nDice Roll : ")
	for i := 0; i < len(dice); i++ {
		fmt.Printf("%d ", dice[i])
	}
	fmt.Println("")
}

// Method to delay execution. Used to make the terminal return easier to read.
func slowDown(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
