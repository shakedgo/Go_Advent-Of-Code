package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	allRe       = regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	operationRe = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	digitsRe    = regexp.MustCompile(`\d+`)
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(input)

	// r := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	matches := allRe.FindAllString(content, -1)

	sumPart1 := 0
	for _, match := range matches {
		sumPart1 += getSum(match)
	}
	fmt.Println("Sum Part 1: ", sumPart1)

	// re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	allMatches := operationRe.FindAllString(content, -1)

	sumPart2 := 0
	isEnabled := true
	for _, match := range allMatches {
		switch match {
		case "do()":
			isEnabled = true
		case "don't()":
			isEnabled = false
		default:
			if isEnabled {
				sumPart2 += getSum(match)
			}
		}
	}

	fmt.Println("Sum Part 2: ", sumPart2)
}

func getSum(s string) int {
	nums := digitsRe.FindAllString(s, -1)

	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])

	return a * b
}
