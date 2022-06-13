package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	var measurements []int

	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		measurements = append(measurements, m)
	}

	file.Close()

	increments := GetIncrements(measurements)

	fmt.Println(increments)
}

func GetIncrements(measurements []int) int {
	increments := 0

	for i := 0; i < len(measurements); i++ {
		if i < len(measurements)-3 {
			sum1 := measurements[i] + measurements[i+1] + measurements[i+2]
			sum2 := measurements[i+1] + measurements[i+2] + measurements[i+3]
			if sum2 > sum1 {
				increments += 1
			}
		}
	}

	return increments
}
