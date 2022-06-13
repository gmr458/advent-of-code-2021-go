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

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var commands []Command

	for scanner.Scan() {
		dataCommand := strings.Split(scanner.Text(), " ")

		direction := dataCommand[0]
		units, err := strconv.Atoi(dataCommand[1])

		if err != nil {
			log.Fatal(err)
		}

		c := Command{direction, units}
		commands = append(commands, c)
	}

	file.Close()

	p := Position{}

	p.executeCommands(commands)
	p.executeMultiplication()

	fmt.Printf("Horizontal = %d\n", p.horizontal)
	fmt.Printf("Depth      = %d\n\n", p.depth)
	fmt.Printf("%d x %d = %d\n", p.horizontal, p.depth, p.multiplication)
}

type Command struct {
	direction string
	units     int
}

type Position struct {
	horizontal     int
	depth          int
	multiplication int
}

func (p *Position) forward(units int) {
	p.horizontal += units
}

func (p *Position) down(units int) {
	p.depth += units
}

func (p *Position) up(units int) {
	p.depth -= units
}

func (p *Position) executeCommands(commands []Command) {
	for _, c := range commands {
		switch c.direction {
		case "forward":
			p.forward(c.units)
		case "down":
			p.down(c.units)
		case "up":
			p.up(c.units)
		}
	}
}

func (p *Position) executeMultiplication() {
	p.multiplication = p.horizontal * p.depth
}
