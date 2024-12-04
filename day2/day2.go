package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkDifferences(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		diff := math.Abs(float64(nums[i+1] - nums[i]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isMonotonic(nums []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			ascending = false
		}
		if nums[i] <= nums[i+1] {
			descending = false
		}

		if !ascending && !descending {
			return false
		}
	}

	return ascending || descending
}

func isSafe(nums []int) bool {
	return isMonotonic(nums) && checkDifferences(nums)
}

func isSafeWithDampener(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		newNums := make([]int, 0)
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

		if isSafe(newNums) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening input.txt: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0
	safeCountPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		numStrings := strings.Fields(line)
		nums := make([]int, len(numStrings))

		for i, numStr := range numStrings {
			nums[i], _ = strconv.Atoi(numStr)
		}

		if isSafe(nums) {
			safeCount++
		}
		if isSafeWithDampener(nums) {
			safeCountPart2++
		}
	}

	fmt.Printf("Part 1: %d\n", safeCount)
	fmt.Printf("Part 2: %d\n", safeCountPart2)
}
