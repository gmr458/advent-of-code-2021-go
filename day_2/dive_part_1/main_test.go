package main

import "testing"

func TestPosition(t *testing.T) {
	commands := []Command{{"forward", 5}, {"down", 5}, {"forward", 8}, {"up", 3}, {"down", 8}, {"forward", 2}}

	p := Position{}

	p.executeCommands(commands)
	p.executeMultiplication()

	wantHorizontal := 15
	wantDepth := 10
	wantMultiplication := 150

	if p.horizontal != wantHorizontal {
		t.Errorf("Want Position.horizontal equal to %d. Got Position.horizontal equal to %d", wantHorizontal, p.horizontal)
	}

	if p.depth != wantDepth {
		t.Errorf("Want Position.depth equal to %d. Got Position.depth equal to %d", wantDepth, p.depth)
	}

	if p.multiplication != wantMultiplication {
		t.Errorf("Want Position.multiplication equal to %d. Got Position.multiplication equal to %d", wantMultiplication, p.multiplication)
	}
}
