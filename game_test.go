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
	pointsValidTests := []struct {
		dice [6]int
		want int
	}{
		{[6]int{1, 2, 5, 2, 3, 4}, 150},
		{[6]int{1, 1, 1, 2, 3, 4}, 1000},
		{[6]int{1, 1, 1, 1, 3, 4}, 1100},
		{[6]int{1, 1, 1, 1, 1, 1}, 2000},
		{[6]int{5, 2, 3, 4, 6, 6}, 50},
		{[6]int{5, 5, 5, 4, 6, 6}, 500},
		{[6]int{5, 5, 5, 5, 5, 5}, 1000},
		{[6]int{2, 2, 2, 3, 3, 4}, 200},
		{[6]int{2, 2, 2, 3, 3, 3}, 500},
		{[6]int{4, 4, 4, 6, 6, 6}, 1000},
		{[6]int{2, 2, 2, 2, 2, 2}, 400},
	}

	for _, tt := range pointsValidTests {
		got := FindPoints(tt.dice)
		if got != tt.want {
			t.Errorf("got %d want %d given %v", got, tt.want, tt.dice)
		}
	}
}
