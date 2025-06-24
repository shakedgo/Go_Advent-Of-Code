package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("controlInput.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(input)

	// var orderRules [][2]int
	var updatePages [][]int
	ruleMap := make(map[int]map[int]bool)
	inSecondSection := false

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			inSecondSection = true
			continue
		}

		if inSecondSection {
			strNums := strings.Split(string(line), ",")
			var page []int
			for _, str := range strNums {
				num, _ := strconv.Atoi(str)
				page = append(page, num)
			}
			updatePages = append(updatePages, page)
		} else {
			rule := strings.Split(string(line), "|")
			num1, _ := strconv.Atoi(rule[0])
			num2, _ := strconv.Atoi(rule[1])
			if ruleMap[num1] == nil {
				ruleMap[num1] = make(map[int]bool)
			}
			ruleMap[num1][num2] = true
		}
	}

	var sumMiddles int

	// Needs Change - Very complex.
	for _, update := range updatePages {
		isValid := true
		seen := make(map[int]bool)
		for _, num := range update {
			if invalids, exists := ruleMap[num]; exists {
				for invalid := range invalids {
					if seen[invalid] {
						isValid = false
						break
					}
				}
			}
			seen[num] = true
		}
		if isValid {
			middleIndex := len(update) / 2
			middleItem := update[middleIndex]
			sumMiddles += middleItem
		}

	}

	// fmt.Println("Rules:", orderRules)
	// fmt.Println("Pages:", updatePages)

	fmt.Println("Part 1 Result:", sumMiddles)
}
