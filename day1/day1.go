package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var part1ans = 0
	var part2ans = 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening input.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
		freq1[num1]++
		freq2[num2]++
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i := 0; i < len(list1); i++ {
		part1ans = part1ans + int(math.Abs(float64(list1[i]-list2[i])))
	}

	fmt.Printf("Part 1: %v\n", part1ans)

	for i := 0; i < len(list1); i++ {
		if value, exists := freq2[list1[i]]; exists {
			part2ans = part2ans + (list1[i] * value)
		}
	}

	fmt.Printf("Part 2: %v\n", part2ans)
}
