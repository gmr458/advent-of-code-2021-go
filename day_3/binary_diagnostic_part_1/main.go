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

	gammaRate := getGammaRate(diagnosticReport)
	epsilonRate := getEpsilonRate(diagnosticReport)

	powerConsumption := getPowerConsumption(gammaRate, epsilonRate)

	fmt.Printf("Gamma rate        = %d\n", gammaRate)
	fmt.Printf("Epsilon rate      = %d\n", epsilonRate)
	fmt.Printf("Power consumption = %d\n", powerConsumption)
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

const (
	gammaRate   = "gammaRate"
	epsilonRate = "epsilonRate"
)

func getMostCommonBit(typeRate string, diagnosticReport [][]string, column int) string {
	var zeros int
	var ones int

	for _, row := range diagnosticReport {
		if row[column] == "0" {
			zeros += 1
			continue
		}

		ones += 1
	}

	if typeRate == gammaRate {
		if zeros > ones {
			return "0"
		}

		return "1"
	}

	if zeros < ones {
		return "0"
	}

	return "1"
}

func getGammaRate(diagnosticReport [][]string) int64 {
	var gammaRateString string

	for column := range diagnosticReport[0] {
		bitGamma := getMostCommonBit(gammaRate, diagnosticReport, column)
		gammaRateString += bitGamma
	}

	gammaRateDecimal, err := strconv.ParseInt(gammaRateString, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	return gammaRateDecimal
}

func getEpsilonRate(diagnosticReport [][]string) int64 {
	var epsilonRateString string

	for column := range diagnosticReport[0] {
		bitEpsilon := getMostCommonBit(epsilonRate, diagnosticReport, column)
		epsilonRateString += bitEpsilon
	}

	epsilonRateDecimal, err := strconv.ParseInt(epsilonRateString, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	return epsilonRateDecimal
}

func getPowerConsumption(gammaRate int64, epsilonRate int64) int64 {
	return gammaRate * epsilonRate
}
