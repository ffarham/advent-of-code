package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() (output []string) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}
	return
}

func getEmptySplit(n int, value string) (output []string) {
	for i := 0; i < n; i++ {
		output = append(output, value)
	}
	return
}

func part1(input []string) (output int) {

	lasers := getEmptySplit(len(input[0]), " ")
	for _, line := range input {
		newLasers := getEmptySplit(len(input[0]), " ")
		for cIndex, char := range line {
			if char == 'S' {
				newLasers[cIndex] = "|"
				break
			}
			if lasers[cIndex] == "|" {
				if char == '^' {
					output++
					if cIndex > 0 {
						newLasers[cIndex-1] = "|"
					}
					if cIndex < len(input[0])-1 {
						newLasers[cIndex+1] = "|"
					}
				} else if char == '.' {
					if lasers[cIndex] == "|" {
						newLasers[cIndex] = "|"
					}
				}
			}
		}
		lasers = newLasers
	}
	return
}

func part2(input []string) (output int) {
	lasers := getEmptySplit(len(input[0]), "0")
	for _, line := range input {
		newLasers := getEmptySplit(len(input[0]), "0")
		for cIndex, char := range line {
			if char == 'S' {
				newLasers[cIndex] = "1"
				break
			}
			if lasers[cIndex] != "0" {
				laserDigit, _ := strconv.Atoi(lasers[cIndex])
				if char == '^' {
					if cIndex > 0 {
						lhsDigit, _ := strconv.Atoi(newLasers[cIndex-1])
						lhsDigit += laserDigit
						lhsDigitStr := strconv.Itoa(lhsDigit)
						newLasers[cIndex-1] = lhsDigitStr
					}
					if cIndex < len(input[0])-1 {
						rhsDigit, _ := strconv.Atoi(newLasers[cIndex+1])
						rhsDigit += laserDigit
						rhsDigitStr := strconv.Itoa(rhsDigit)
						newLasers[cIndex+1] = rhsDigitStr
					}
				} else if char == '.' {
					if lasers[cIndex] != "0" {
						prevDigit, _ := strconv.Atoi(lasers[cIndex])
						currDigit, _ := strconv.Atoi(newLasers[cIndex])
						newDigit := strconv.Itoa(prevDigit + currDigit)
						newLasers[cIndex] = newDigit
					}
				}
			}
		}
		lasers = newLasers
	}

	for _, laser := range lasers {
		digit, _ := strconv.Atoi(laser)
		output += digit
	}
	return
}

func main() {
	input := getInput()
	fmt.Printf("[Part 1] Ans: %v\n", part1(input))
	fmt.Printf("[Part 2] Ans: %v\n", part2(input))
}
