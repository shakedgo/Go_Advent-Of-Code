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
	var sumFixedMiddles int

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
		} else {
			sorted := topoSort(update, ruleMap)
			if len(sorted) > 0 {
				middle := sorted[len(sorted)/2]
				sumFixedMiddles += middle
			}
		}

	}

	fmt.Println("Part 1 Result:", sumMiddles)
	fmt.Println("Part 2 Result:", sumFixedMiddles)
}

func topoSort(updates []int, rules map[int]map[int]bool) []int {
	inDegree := make(map[int]int)
	adj := make(map[int][]int)

	update := make(map[int]bool)
	for _, n := range updates {
		update[n] = true
	}

	for from := range update {
		for to := range update {
			if rules[from] != nil && rules[from][to] {
				adj[to] = append(adj[to], from)
				inDegree[from]++
			}
		}
	}

	var queue []int
	for _, node := range updates {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sorted = append(sorted, curr)
		for _, neighbor := range adj[curr] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(sorted) != len(updates) {
		return nil
	}

	return sorted
}
