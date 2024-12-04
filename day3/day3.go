package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening input.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	sumPart1 := processInput(input)
	fmt.Printf("Part 1: %d\n", sumPart1)

	sumPart2 := processInputPart2(input)
	fmt.Printf("Part 2: %d\n", sumPart2)
}

func processInput(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		product := x * y
		sum += product
	}

	return sum
}

type Operation struct {
	start    int
	opType   string
	x, y     int
	disabled bool
}

func processInputPart2(input string) int {
	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	var operations []Operation

	mulMatches := mulRe.FindAllStringSubmatchIndex(input, -1)
	for _, match := range mulMatches {
		x, _ := strconv.Atoi(input[match[2]:match[3]])
		y, _ := strconv.Atoi(input[match[4]:match[5]])
		operations = append(operations, Operation{
			start:  match[0],
			opType: "mul",
			x:      x,
			y:      y,
		})
	}

	doMatches := doRe.FindAllStringIndex(input, -1)
	for _, match := range doMatches {
		operations = append(operations, Operation{
			start:  match[0],
			opType: "do",
		})
	}

	dontMatches := dontRe.FindAllStringIndex(input, -1)
	for _, match := range dontMatches {
		operations = append(operations, Operation{
			start:  match[0],
			opType: "don't",
		})
	}
	sortOperations(operations)

	enabled := true
	sum := 0

	for i := range operations {
		switch operations[i].opType {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			if enabled {
				product := operations[i].x * operations[i].y
				sum += product
			}
		}
	}

	return sum
}

func sortOperations(ops []Operation) {
	n := len(ops)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ops[j].start > ops[j+1].start {
				ops[j], ops[j+1] = ops[j+1], ops[j]
			}
		}
	}
}
