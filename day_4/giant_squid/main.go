package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pass the input file as argument.")
		os.Exit(0)
	}

	filename := os.Args[1]

	nums := GetNumbers(filename)

	boards := GetBoards(filename)

	boardWinner, boardLastWinner := GetWinBoard(nums, boards)

	fmt.Println("Winner Board:")
	fmt.Printf(
		"Id: %d, Attempts: %d, Win Number: %d, Sum Unmarked Numbers: %d, Final Score: %d\n",
		boardWinner.Id,
		boardWinner.Attempts,
		boardWinner.WinNumber,
		boardWinner.SumUnmarkedNumbers,
		boardWinner.FinalScore,
	)

	fmt.Println("\nLast Winner Board:")
	fmt.Printf(
		"Id: %d, Attempts: %d, Win Number: %d, Sum Unmarked Numbers: %d, Final Score: %d\n",
		boardLastWinner.Id,
		boardLastWinner.Attempts,
		boardLastWinner.WinNumber,
		boardLastWinner.SumUnmarkedNumbers,
		boardLastWinner.FinalScore,
	)
}

func GetNumbers(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums []int

	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, ",")

		for _, s := range numsStr {
			n, err := strconv.Atoi(s)

			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		break
	}

	file.Close()

	return nums
}

func GetBoards(filename string) []*Board {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var boards []*Board

	actualLine := 1
	idBoards := 1

	for scanner.Scan() {
		if actualLine > 2 {
			b := Board{Id: idBoards}

			firstNumsString := strings.Split(scanner.Text(), " ")
			firstNumsInt := StringSliceToIntSlice(firstNumsString)

			b.Grid = append(b.Grid, firstNumsInt)

			for scanner.Scan() && scanner.Text() != "" {

				numsString := strings.Split(scanner.Text(), " ")
				numsInt := StringSliceToIntSlice(numsString)

				b.Grid = append(b.Grid, numsInt)
			}

			boards = append(boards, &b)
			idBoards += 1
		}

		actualLine += 1
	}

	file.Close()

	return boards
}

func StringSliceToIntSlice(numsString []string) []int {
	var numsInt []int

	for _, s := range numsString {
		if s == "" {
			continue
		}

		n, err := strconv.Atoi(s)

		if err != nil {
			log.Fatal(err)
		}

		numsInt = append(numsInt, n)
	}

	return numsInt
}

type Board struct {
	Id                 int
	Grid               [][]int
	Attempts           int
	WinNumber          int
	SumUnmarkedNumbers int
	FinalScore         int
}

func (b *Board) CheckRowsAndColls(num int) bool {
	for row := range b.Grid[0] {
		win := false

		for col := range b.Grid {
			if col < len(b.Grid)-1 {
				win = b.Grid[row][col] == b.Grid[row][col+1]

				if !win {
					break
				}
			}
		}

		if win {
			b.WinNumber = num
			return true
		}
	}

	for col := range b.Grid {
		win := false

		for row := range b.Grid[col] {
			if row < len(b.Grid[row])-1 {
				win = b.Grid[row][col] == b.Grid[row+1][col]

				if !win {
					break
				}
			}
		}

		if win {
			b.WinNumber = num
			return true
		}
	}

	return false
}

func (b *Board) MarkAndCheckPossibleVictory(num int) bool {
	for row := range b.Grid {
		for col := range b.Grid[row] {
			if b.Grid[row][col] == num {
				b.Grid[row][col] = -1
			}
		}
	}

	b.Attempts = b.Attempts + 1

	return b.CheckRowsAndColls(num)
}

func (b *Board) SetSumUnmarkedNumbers() {
	for row := range b.Grid {
		for col := range b.Grid[row] {
			if b.Grid[row][col] != -1 {
				b.SumUnmarkedNumbers += b.Grid[row][col]
			}
		}
	}
}

func (b *Board) SetFinalScore() {
	b.FinalScore = b.SumUnmarkedNumbers * b.WinNumber
}

func GetWinBoard(nums []int, boards []*Board) (Board, Board) {
	for _, b := range boards {
		for _, n := range nums {
			win := b.MarkAndCheckPossibleVictory(n)

			if win {
				break
			}
		}

		b.SetSumUnmarkedNumbers()
		b.SetFinalScore()
	}

	boardWinner := boards[0]
	boardLastWinner := boards[0]

	for _, b := range boards {
		if b.Attempts < boardWinner.Attempts {
			boardWinner = b
		}
	}

	for _, b := range boards {
		if b.Attempts > boardLastWinner.Attempts {
			boardLastWinner = b
		}
	}

	return *boardWinner, *boardLastWinner
}
