package dice

import (
	"testing"
)

func TestRoll(t *testing.T) {
	// create zeroed dice array
	dice := [6]int{}

	// populate with random number between 1 & 6 to simulate dice roll
	newDice := Roll(dice)

	// test if diceRoll[0] is between 1 & 6
	got := newDice[0]
	want := false

	for _, num := range newDice {
		got = newDice[num]

		if got > 0 && got < 7 {
			want = true
		} else {
			break
		}
	}

	if want == false {
		t.Errorf("got %d which exceeds dice value range (1-6) given, %v", got, newDice)
	}
}
