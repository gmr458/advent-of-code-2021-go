package main

import "testing"

func TestGetIncrements(t *testing.T) {
	measurements := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	increments := GetIncrements(measurements)

	wantIncrements := 7

	if increments != wantIncrements {
		t.Errorf("Want increments equal to %d. Got increments equal to %d", wantIncrements, increments)
	}
}
