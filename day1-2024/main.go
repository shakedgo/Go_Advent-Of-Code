package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	pt, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(pt)))
	var s1 []int
	var s2 []int
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), "   ")
		num, _ := strconv.Atoi(words[0])
		num2, _ := strconv.Atoi(words[1])
		s1 = append(s1, num)
		s2 = append(s2, num2)
	}

	partTwo := secondPart(s1, s2)
	totalDiff := 0
	// I use s1 length because it is said that they're the same length.
	for len(s1) > 0 {
		minIdx1 := 0
		minIdx2 := 0
		for i := 0; i < len(s1); i++ {
			if s1[i] < s1[minIdx1] {
				minIdx1 = i
			}
			if s2[i] < s2[minIdx2] {
				minIdx2 = i
			}
		}

		totalDiff += int(math.Abs(float64(s1[minIdx1] - s2[minIdx2])))

		s1 = append(s1[:minIdx1], s1[minIdx1+1:]...)
		s2 = append(s2[:minIdx2], s2[minIdx2+1:]...)
	}

	fmt.Println("TotalDiff: ", totalDiff)
	fmt.Println("Part 2: ", partTwo)
}

func secondPart(s1, s2 []int) int {
	var similarityCount int

	for i := 0; i < len(s1); i++ {
		num := s1[i]
		repeats := 0
		for l := 0; l < len(s2); l++ {
			if s2[l] == num {
				repeats++
			}
		}
		similarityCount += num * repeats
	}

	return similarityCount
}
