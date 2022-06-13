package main

import "testing"

var nums = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var b = Board{
	Grid: [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	},
}

func TestMarkAndCheckPossibleVictory(t *testing.T) {
	for _, v := range nums {
		win := b.MarkAndCheckPossibleVictory(v)

		if win {
			break
		}
	}

	b.SetSumUnmarkedNumbers()
	b.SetFinalScore()

	wantAttempts := 12

	if b.Attempts != wantAttempts {
		t.Errorf("Want Board.Attempts == %d. Got Board.Attempts == %d", wantAttempts, b.Attempts)
	}

	wantWinNumber := 24

	if b.WinNumber != wantWinNumber {
		t.Errorf("Want Board.WinNumber == %d. Got Board.WinNumber == %d", wantWinNumber, b.WinNumber)
	}

	wantSumUnnmarkedNumbers := 188

	if b.SumUnmarkedNumbers != wantSumUnnmarkedNumbers {
		t.Errorf("Want Board.WinNumber == %d. Got Board.WinNumber == %d", wantSumUnnmarkedNumbers, b.SumUnmarkedNumbers)
	}

	wantFinalScore := 4512

	if b.FinalScore != wantFinalScore {
		t.Errorf("Want Board.FinalScore == %d. Got Board.FinalScore == %d", wantFinalScore, b.FinalScore)
	}
}
