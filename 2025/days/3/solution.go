package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func getInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()

	output := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return output
}

func part1(input []string) (totalJoltage int) {
	for _, line := range input {
		lPtr, rPtr := len(line)-2, len(line)-1
		for i := lPtr - 1; i >= 0; i-- {
			if line[i] >= line[lPtr] {
				lPtr = i
			}
		}

		for j := rPtr - 1; j > lPtr; j-- {
			if line[j] >= line[rPtr] {
				rPtr = j
			}
		}

		joltage, _ := strconv.Atoi(string(line[lPtr]) + string(line[rPtr]))
		totalJoltage += joltage
	}
	return totalJoltage
}

// Time Complexity: O(n) where n in the length of the sequence

func part2(input []string) (output int) {

	for _, line := range input {

		// build a map of digit and the indices of its occurences
		store := make(map[int]*list.List)
		for i := 0; i <= 9; i++ {
			store[i] = list.New()
		}
		for index, char := range line {
			digit, _ := strconv.Atoi(string(char))
			store[digit].PushBack(index)
		}

		maxStr := ""

		farLeft := -1
		changeFlag := false
		for i := len(line) - 12; i < len(line); i++ {
			digit, _ := strconv.Atoi(string(line[i]))
		Digit:
			for j := 9; j >= digit; j-- {
				for store[j].Front() != nil {
					front := store[j].Front()
					newIndex := front.Value.(int)
					store[j].Remove(front)
					if newIndex > farLeft && newIndex < i {
						maxStr += strconv.Itoa(j)
						farLeft = newIndex
						changeFlag = true
						break Digit
					}
				}
			}
			if !changeFlag {
				maxStr += string(line[i:])
				break
			}
			changeFlag = false
		}

		maxInt, _ := strconv.Atoi(maxStr)
		output += maxInt
	}

	return
}

func main() {
	input := getInput()

	fmt.Printf("[Part 1] Total Joltage is %v\n", part1(input))
	fmt.Printf("[Part 2] Total Joltage is %v\n", part2(input))
}
