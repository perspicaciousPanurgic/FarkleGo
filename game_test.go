package main

import "testing"

func TestisValidNumPlayers(bool, t *testing.T) {
	got := isValidNumPlayers(5)
	want := false

	if got != false {
		t.Errorf("got %d want %d given 5", got, want)
	}
}
