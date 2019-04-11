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
		newDice := Roll(&dice, &numDice)

		// test if diceRoll[0] is between 1 & 6
		got := len(newDice)
		want := 6

		if got != want {
			t.Errorf("got %d expect %d given, %v", got, want, newDice)
		}
	})

	t.Run("values", func(t *testing.T) {
		// populate with random number between 1 & 6 to simulate dice roll
		newDice := Roll(&dice, &numDice)

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

func TestCountDice(t *testing.T) {
	isCountValid := []struct {
		dice  [6]int
		count [6]int
	}{
		{[6]int{1, 2, 3, 4, 5, 6}, [6]int{1, 1, 1, 1, 1, 1}},
		{[6]int{6, 5, 4, 3, 2, 1}, [6]int{1, 1, 1, 1, 1, 1}},
		{[6]int{1, 1, 1, 1, 1, 1}, [6]int{6, 0, 0, 0, 0, 0}},
		{[6]int{6, 6, 6, 6, 6, 6}, [6]int{0, 0, 0, 0, 0, 6}},
	}

	for _, tt := range isCountValid {
		got := CountDice(&tt.dice)
		if got != tt.count {
			t.Errorf("got %v want %v given %v", got, tt.count, tt.dice)
		}
	}
}
