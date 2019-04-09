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
		dice    [6]int
		points  int
		numDice int
	}{
		{[6]int{1, 2, 5, 2, 3, 4}, 150, 2},
		{[6]int{1, 1, 1, 2, 3, 4}, 1000, 3},
		{[6]int{1, 1, 1, 1, 3, 4}, 1100, 4},
		{[6]int{1, 1, 1, 1, 1, 1}, 2000, 6},
		{[6]int{5, 2, 3, 4, 6, 6}, 50, 1},
		{[6]int{5, 5, 5, 4, 6, 6}, 500, 3},
		{[6]int{5, 5, 5, 5, 5, 5}, 1000, 6},
		{[6]int{2, 2, 2, 3, 3, 4}, 200, 3},
		{[6]int{2, 2, 2, 3, 3, 3}, 500, 6},
		{[6]int{4, 4, 4, 6, 6, 6}, 1000, 6},
		{[6]int{2, 2, 2, 2, 2, 2}, 400, 6},
	}

	for _, tt := range pointsValidTests {
		got := FindPoints(tt.dice)
		if got.points != tt.points {
			t.Errorf("Points: got %d want %d given %v", got.points, tt.points, tt.dice)
		}
		if got.numDice != tt.numDice {
			t.Errorf("numDice: got %d want %d given %v", got.numDice, tt.numDice, tt.dice)
		}
	}
}

func TestKeepPoints(t *testing.T) {
	keepPointsTests := []struct {
		points int
		want   int
	}{
		{100, 100},
		{100, 200},
		{1000, 1200},
	}

	player := Player{1, "Adam", 0}

	for _, tt := range keepPointsTests {
		KeepPoints(tt.points, &player)
		got := player.score
		if got != tt.want {
			t.Errorf("got %d want %d", got, tt.want)
		}
	}
}

func TestEndRound(t *testing.T) {
	roundTest := []struct {
		points  int
		numDice int
		score   int
		farkle  bool
	}{
		{0, 0, 0, true},
		{200, 2, 0, false},
		{800, 4, 0, false},
		{0, 2, 1000, true},
		{800, 2, 0, false},
		{1000, 2, 0, true},
		{1000, 6, 0, false},
		{400, 3, 1000, true},
		{400, 6, 1000, false},
		{0, 6, 0, true},
		{100, 6, 0, false},
	}

	for _, tt := range roundTest {
		farkle := EndRound(tt.points, tt.numDice, tt.score)
		if farkle != tt.farkle {
			t.Errorf("got %t want %t given %d points, %d dice used and %d score", farkle, tt.farkle, tt.points, tt.numDice, tt.score)
		}
	}
}
