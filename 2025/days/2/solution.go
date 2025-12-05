package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (output string) {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output += scanner.Text()
	}
	return
}

func part1(input string) (output int) {
	ranges := strings.Split(input, ",")
	for _, currRange := range ranges {
		nums := strings.Split(currRange, "-")
		startStr, endStr := nums[0], nums[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		for i := start; i <= end; i++ {
			iStr := strconv.Itoa(i)
			n := len(iStr)
			if n%2 != 0 {
				continue
			}
			mid := n / 2
			firstHalf, secondHalf := iStr[:mid], iStr[mid:]
			if firstHalf == secondHalf {
				output += i
			}
		}
	}
	return
}

func containsRepeats(numStr string) bool {

	n := len(numStr)
	windowSize := 1
	for windowSize <= n/2 {
		if n%windowSize != 0 {
			windowSize++
			continue
		}

		prevWindow := ""
		isRepeat := true
		for i := 0; i < len(numStr); i += windowSize {
			window := numStr[i : i+windowSize]
			if prevWindow == "" {
				prevWindow = window
				continue
			}
			if prevWindow != window {
				isRepeat = false
				break
			}
		}
		if isRepeat {
			return true
		}
		windowSize++
	}
	return false
}

func part2(input string) (output int) {
	ranges := strings.Split(input, ",")
	for _, currRange := range ranges {
		nums := strings.Split(currRange, "-")
		startStr, endStr := nums[0], nums[1]
		start, _ := strconv.Atoi(startStr)
		end, _ := strconv.Atoi(endStr)
		for i := start; i <= end; i++ {
			iStr := strconv.Itoa(i)
			if containsRepeats(iStr) {
				output += i
			}
		}
	}
	return
}

func main() {
	input := getInput()

	fmt.Printf("[Part 1] Ans: %v\n", part1(input))
	fmt.Printf("[Part 2] Ans: %v\n", part2(input))
}
