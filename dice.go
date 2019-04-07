package main

import "math/rand"
import "time"

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
