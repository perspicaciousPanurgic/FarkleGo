package main

import (
	"testing"
)

func TestRoll(t *testing.T) {
	// create zeroed dice array
	dice := [6]int{}
	numDice := 3

	t.Run("length", func(t *testing.T) {
		// populate with random number between 1 & 6 to simulate dice roll
		newDice := Roll(dice, numDice)

		// test if diceRoll[0] is between 1 & 6
		got := len(newDice)
		want := 6

		if got != want {
			t.Errorf("got %d expect %d given, %v", got, want, newDice)
		}
	})

	t.Run("values", func(t *testing.T) {
		// populate with random number between 1 & 6 to simulate dice roll
		newDice := Roll(dice, numDice)

		// test if diceRoll[0] is between 1 & 6
		got := newDice[0]
		want := false

		for i := 0; i < numDice; i++ {
			got = newDice[i]

			if got > 0 && got < 7 {
				want = true
			} else {
				break
			}
		}

		if want == false {
			t.Errorf("got %d which exceeds dice value range (1-6) given, %v", got, newDice)
		}
	})
}
