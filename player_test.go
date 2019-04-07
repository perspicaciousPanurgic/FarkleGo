package main

import "testing"

func TestChangeName(t *testing.T) {
	// Initialize player
	playerOne := Player{0, "Tom", 0}

	// Change player name
	newName := "Bilbo"
	playerOne.ChangeName(newName)

	got := playerOne.name
	want := newName

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestAddScore(t *testing.T) {
	// Initialize player
	playerOne := Player{0, "Tom", 0}

	// Add score
	points := 100
	playerOne.AddScore(100)

	got := playerOne.score
	want := points

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
