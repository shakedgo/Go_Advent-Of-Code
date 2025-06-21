package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [][2]int{
	{1, 1}, {-1, -1}, {-1, 1}, {1, -1},
	{0, -1}, {0, 1}, {1, 0}, {-1, 0},
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(input)

	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Find XMAS
	countRepeats1 := 0
	countRepeats2 := 0
	part1Word := "XMAS"
	part2Word := "MAS"
	part2WordRev := reverse(part2Word)
	for lineIndex, line := range lines {
		letters := strings.Split(line, "")
		for letterIndex := range letters {
			for _, dir := range directions {
				if checkRepeats(part1Word, lines, lineIndex, letterIndex, dir[0], dir[1]) {
					countRepeats1++
				}
			}

			if checkRepeats(part2Word, lines, lineIndex, letterIndex, 1, 1) &&
				checkRepeats(part2WordRev, lines, lineIndex, letterIndex+2, 1, -1) {
				// M . S
				// . A .
				// M . S
				countRepeats2++
			}
			if checkRepeats(part2Word, lines, lineIndex, letterIndex, 1, -1) &&
				checkRepeats(part2WordRev, lines, lineIndex, letterIndex-2, 1, 1) {
				// S . M
				// . A .
				// S . M
				countRepeats2++
			}
			if checkRepeats(part2Word, lines, lineIndex, letterIndex, 1, 1) &&
				checkRepeats(part2Word, lines, lineIndex, letterIndex+2, 1, -1) {
				// M . M
				// . A .
				// S . S
				countRepeats2++
			}
			if checkRepeats(part2WordRev, lines, lineIndex, letterIndex, 1, 1) &&
				checkRepeats(part2WordRev, lines, lineIndex, letterIndex+2, 1, -1) {
				// S . S
				// . A .
				// M . M
				countRepeats2++
			}
		}
	}

	fmt.Println("Repeats Part 1: ", countRepeats1)
	fmt.Println("Repeats Part 2: ", countRepeats2)
}

func checkRepeats(word string, lines []string, row, col, dRow, dCol int) bool {
	for i := 0; i < len(word); i++ {
		r, c := row+i*dRow, col+i*dCol
		if r < 0 || r >= len(lines) || c < 0 || c >= len(lines[r]) {
			return false
		}
		if lines[r][c] != word[i] {
			return false
		}
	}
	return true
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
