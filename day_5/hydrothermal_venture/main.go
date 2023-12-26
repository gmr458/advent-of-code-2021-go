package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start *Point
	End   *Point
}

func solutionPart1(lines []Line) int {
	register := make(map[string]int)

	for _, line := range lines {
		// vertical line
		if line.Start.X == line.End.X {
			var start, end int
			if line.Start.Y > line.End.Y {
				start = line.End.Y
				end = line.Start.Y
			} else {
				start = line.Start.Y
				end = line.End.Y
			}

			for i := start; i <= end; i++ {
				register[fmt.Sprintf("%d,%d", line.Start.X, i)] += 1
			}
		}

		// horizontal line
		if line.Start.Y == line.End.Y {
			var start, end int
			if line.Start.X > line.End.X {
				start = line.End.X
				end = line.Start.X
			} else {
				start = line.Start.X
				end = line.End.X
			}

			for i := start; i <= end; i++ {
				register[fmt.Sprintf("%d,%d", i, line.Start.Y)] += 1
			}
		}
	}

	var count int

	for _, v := range register {
		if v >= 2 {
			count++
		}
	}

	return count
}

// NOT FINISHED
func solutionPart2(lines []Line) int {
	register := make(map[string]int)

	for _, line := range lines {
		// vertical line
		if line.Start.X == line.End.X {
			var start, end int
			if line.Start.Y > line.End.Y {
				start = line.End.Y
				end = line.Start.Y
			} else {
				start = line.Start.Y
				end = line.End.Y
			}

			for i := start; i <= end; i++ {
				register[fmt.Sprintf("%d,%d", line.Start.X, i)] += 1
			}

			continue
		}

		// horizontal line
		if line.Start.Y == line.End.Y {
			var start, end int
			if line.Start.X > line.End.X {
				start = line.End.X
				end = line.Start.X
			} else {
				start = line.Start.X
				end = line.End.X
			}

			for i := start; i <= end; i++ {
				register[fmt.Sprintf("%d,%d", i, line.Start.Y)] += 1
			}

			continue
		}

		// diagonal line
		if line.Start.X != line.End.X || line.Start.Y != line.End.Y {
			//
			// var xStart, xEnd, yStart, yEnd int
			var xStart, xEnd, yStart int

			//            8 > 0
			if line.Start.X > line.End.X {
				// 0
				xStart = line.End.X

				// 8
				xEnd = line.Start.X
			} else {
				xStart = line.Start.X
				xEnd = line.End.X
			}

			yStart = xEnd
			// yEnd = xStart

			for xStart <= xEnd {
				register[fmt.Sprintf("%d,%d", xStart, yStart)] += 1
				xStart++
				yStart--
			}
		}
	}

	var count int

	for k, v := range register {
		fmt.Printf("%s = %d\n", k, v)

		if v >= 2 {
			count++
		}
	}

	return count
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pass the input file as argument.")
		os.Exit(0)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []Line

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		if len(line) != 2 {
			panic("a line should hava 2 points, start and end")
		}

		start := line[0]
		end := line[1]

		x1y1 := strings.Split(start, ",")
		if len(x1y1) != 2 {
			panic("a point should hava 2 values, x and y")
		}

		x1, err := strconv.Atoi(x1y1[0])
		if err != nil {
			log.Fatalf("error parsing x1: %s", err.Error())
		}
		y1, err := strconv.Atoi(x1y1[1])
		if err != nil {
			log.Fatalf("error parsing y1: %s", err.Error())
		}

		x2y2 := strings.Split(end, ",")
		if len(x1y1) != 2 {
			panic("a point should hava 2 values, x and y")
		}

		x2, err := strconv.Atoi(x2y2[0])
		if err != nil {
			log.Fatalf("error parsing x2: %s", err.Error())
		}
		y2, err := strconv.Atoi(x2y2[1])
		if err != nil {
			log.Fatalf("error parsing y2: %s", err.Error())
		}

		l := Line{
			Start: &Point{X: x1, Y: y1},
			End:   &Point{X: x2, Y: y2},
		}

		lines = append(lines, l)
	}

	file.Close()

	count := solutionPart1(lines)

	fmt.Println(count)
}
