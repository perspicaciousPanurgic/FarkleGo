package main

import "testing"

func TestNextActivePlayer(t *testing.T) {
	currentPlayer := 1
	numPlayers := 2
	got := NextActivePlayer(currentPlayer, numPlayers)
	want := 0

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestIsValidNumPlayers(t *testing.T) {

	isValidTests := []struct {
		num  int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, true},
		{5, false},
	}

	for _, tt := range isValidTests {
		got := IsValidNumPlayers(tt.num)
		if got != tt.want {
			t.Errorf("got %t want %t given %d", got, tt.want, tt.num)
		}
	}
}

func TestNamePlayer(t *testing.T) {
	name := "Bilbo"
	index := 0
	playerList := []Player{
		Player{1, "Tom", 0},
		Player{2, "Bill", 0},
	}

	NamePlayer(index, name, playerList)

	got := playerList[index+1].name
	want := "Bill"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestFindPoints(t *testing.T) {
	diceRoll := [6]int{1, 1, 2, 2, 3, 4}

	got := FindPoints(diceRoll)
	want := 200

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, diceRoll)
	}
}
