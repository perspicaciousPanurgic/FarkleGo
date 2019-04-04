package main

import "testing"

func TestGreetings(t *testing.T) {
	got := Greetings()
	want := "Let's Play Farkle!"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
