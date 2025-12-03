package main

import (
	"bufio"
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

func part1(input []string) (output int) {
	currentPos := 50
	for _, line := range input {
		ticksStr := line[1:]
		ticks, _ := strconv.Atoi(ticksStr)
		if line[0] == 'L' {
			currentPos -= ticks
		} else {
			currentPos += ticks
		}
		currentPos %= 100
		if currentPos == 0 {
			output++
		}
	}
	return
}

func part2(input []string) (output int) {
	currentPos := 50
	for _, line := range input {
		ticksStr := line[1:]
		ticks, _ := strconv.Atoi(ticksStr)
		if line[0] == 'L' {
			ticksToZero := currentPos
			if ticks >= ticksToZero {
				ticks -= ticksToZero
				currentPos = 0
				if ticksToZero != 0 {
					output++
				}
				output += (ticks / 100)

				currentPos = (100 - (ticks % 100)) % 100

			} else {
				currentPos -= ticks
			}
		} else {
			ticksToZero := 100 - currentPos
			if ticks >= ticksToZero {
				ticks -= ticksToZero
				currentPos = 0
				if ticksToZero != 0 {
					output++
				}
				output += (ticks / 100)
				currentPos = ticks % 100
			} else {
				currentPos += ticks
			}
		}
	}

	return output
}

func main() {
	input := getInput()

	part1Ans := part1(input)
	fmt.Printf("Part 1: %d\n", part1Ans)

	part2Ans := part2(input)
	fmt.Printf("Part 2: %d\n", part2Ans)
}
