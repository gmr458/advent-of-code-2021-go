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

	diagnosticReport := getDiagnosticReport(filename)

	oxygenGeneratorRating := getOxygenGeneratorRating(diagnosticReport, 0)

	CO2ScrubberRating := getCO2ScrubberRating(diagnosticReport, 0)

	lifeSupportRating := getLifeSupportRating(oxygenGeneratorRating, CO2ScrubberRating)

	fmt.Printf("Oxygen Generator Rating = %d\n", oxygenGeneratorRating)
	fmt.Printf("CO2 Scrubber Rating     = %d\n", CO2ScrubberRating)
	fmt.Printf("%d x %d = %d\n", oxygenGeneratorRating, CO2ScrubberRating, lifeSupportRating)
}

func getDiagnosticReport(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var diagnosticReport [][]string

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		diagnosticReport = append(diagnosticReport, line)
	}

	file.Close()

	return diagnosticReport
}

func getOxygenGeneratorRating(diagnosticReport [][]string, column int) int64 {
	var zeros int
	var ones int

	var filtered [][]string

	for _, row := range diagnosticReport {
		if row[column] == "0" {
			zeros += 1
			continue
		}

		ones += 1
	}

	if ones > zeros || ones == zeros {
		for _, row := range diagnosticReport {
			if row[column] == "1" {
				filtered = append(filtered, row)
			}
		}
	} else {
		for _, row := range diagnosticReport {
			if row[column] == "0" {
				filtered = append(filtered, row)
			}
		}
	}

	if len(filtered) == 1 {
		var binary string

		for _, digit := range filtered[0] {
			binary += digit
		}

		decimal, err := strconv.ParseInt(binary, 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		return decimal
	}

	return getOxygenGeneratorRating(filtered, column+1)
}

func getCO2ScrubberRating(diagnosticReport [][]string, column int) int64 {
	var zeros int
	var ones int

	var filtered [][]string

	for _, row := range diagnosticReport {
		if row[column] == "0" {
			zeros += 1
			continue
		}

		ones += 1
	}

	if zeros < ones || zeros == ones {
		for _, row := range diagnosticReport {
			if row[column] == "0" {
				filtered = append(filtered, row)
			}
		}
	} else {
		for _, row := range diagnosticReport {
			if row[column] == "1" {
				filtered = append(filtered, row)
			}
		}
	}

	if len(filtered) == 1 {
		var binary string

		for _, digit := range filtered[0] {
			binary += digit
		}

		decimal, err := strconv.ParseInt(binary, 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		return decimal
	}

	return getCO2ScrubberRating(filtered, column+1)
}

func getLifeSupportRating(oxygenGeneratorRatingBinary int64, CO2ScrubberRatingBinary int64) int64 {
	return oxygenGeneratorRatingBinary * CO2ScrubberRatingBinary
}
