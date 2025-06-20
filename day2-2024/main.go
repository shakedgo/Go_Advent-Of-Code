package main

import (
	"bufio"
	"fmt"
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
	var reports [][]int
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		nunString := strings.Split(scanner.Text(), " ")
		var nums []int
		for i := 0; i < len(nunString); i++ {
			num, _ := strconv.Atoi(nunString[i])
			nums = append(nums, num)
		}
		reports = append(reports, nums)
	}

	// Change should be at least 1 and atmost 3
	// Should be consistent = increasing/decreasing
	validReports := 0
	for _, report := range reports {
		isIncreasing := false
		isValid := true
		for i, el := range report {
			if i+1 < len(report) {
				if i == 0 {
					// First Loop - detect increase or decrease
					if el < report[i+1] {
						isIncreasing = true
						if report[i+1]-el < 1 || report[i+1]-el > 3 {
							isValid = false
						}
					} else {
						isIncreasing = false
						if el-report[i+1] < 1 || el-report[i+1] > 3 {
							isValid = false
						}
					}
				} else if el < report[i+1] && isIncreasing {
					// if element is smaller than next element and is increasing
					if report[i+1]-el < 1 || report[i+1]-el > 3 {
						// check if next element minus element is less than one or more than 3
						isValid = false
					}
				} else if el > report[i+1] && !isIncreasing {
					// if element is bigger than next element and is decreasing
					if el-report[i+1] < 1 || el-report[i+1] > 3 {
						// check if element minus next element is less than one or more than 3
						isValid = false
					}
				} else {
					isValid = false
				}
			}
		}
		if isValid {
			validReports++
		}
	}

	fmt.Println(validReports)

}
