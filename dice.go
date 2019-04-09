package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct to hold roll information for FindPoints()
type rollResult struct {
	points  int
	numDice int
}

//Roll to generate random numbers between 1-6
func Roll(dice [6]int, numDice int) [6]int {

	// generate random source seed using time
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// fill array with zeroes to ensure clean start
	for i := range dice {
		dice[i] = 0
	}

	// start at index 0 and fill array with random values between 1 & 6 until numDice index
	// this simulates rolling numDice number of dice
	for i := 0; i < numDice; i++ {
		dice[i] = r1.Intn(5) + 1
	}

	return dice
}

// Returns [6]int array which contains the number of times each value occurs given a dice roll
func CountDice(dice [6]int) [6]int {
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

func PrintActiveDice(dice [6]int) {
	fmt.Printf("Dice Roll : %v\n", dice)
}

func PrintHeldDice(dice [6]int) {
	fmt.Printf("Held Dice : %v\n", dice)
}
