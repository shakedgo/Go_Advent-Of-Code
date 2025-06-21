package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(input)

	var orderRules [][2]int
	var updatePages [][]int
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
			orderRules = append(orderRules, [2]int{num1, num2})
		}
	}

	var sumMiddles int

	// Needs Change - Very complex.
	for _, page := range updatePages {
		isValid := true
		for numIndex, num := range page {
			for _, ruleNums := range orderRules {
				if ruleNums[0] == num {
					for i := numIndex; i >= 0; i-- {
						if page[i] == ruleNums[1] {
							isValid = false
						}
					}
				}
			}
		}
		if isValid {
			middleIndex := len(page) / 2
			middleItem := page[middleIndex]
			sumMiddles += middleItem
		}

	}

	fmt.Println("Part 1 Result:", sumMiddles)
}
