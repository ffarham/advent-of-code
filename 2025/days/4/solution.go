package main

import (
	"bufio"
	"fmt"
	"os"
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

func getNeighbours(x, y, maxX, maxY int) (output [][]int) {
	if x > 0 && y > 0 {
		output = append(output, []int{x - 1, y - 1})
	}
	if y > 0 {
		output = append(output, []int{x, y - 1})
	}
	if x < maxX && y > 0 {
		output = append(output, []int{x + 1, y - 1})
	}
	if x < maxX {
		output = append(output, []int{x + 1, y})
	}
	if x < maxX && y < maxY {
		output = append(output, []int{x + 1, y + 1})
	}
	if y < maxY {
		output = append(output, []int{x, y + 1})
	}
	if x > 0 && y < maxY {
		output = append(output, []int{x - 1, y + 1})
	}
	if x > 0 {
		output = append(output, []int{x - 1, y})
	}
	return
}

func getCount(neighbours [][]int, input []string) (output int) {
	for _, neighbour := range neighbours {
		nx, ny := neighbour[0], neighbour[1]
		if input[ny][nx] == '@' {
			output++
		}
	}
	return
}

// Time complexity: O(n) where n is the number of data points in the grid.

func part1(input []string) (output int) {
	maxX, maxY := len(input)-1, len(input[0])-1

	for y, line := range input {
		for x, char := range line {
			if char == '.' {
				continue
			}
			neighbours := getNeighbours(x, y, maxX, maxY)
			if getCount(neighbours, input) < 4 {
				output++
			}
		}
	}

	return
}

// Time complexity: O(n^2) where n is the number of data points in the grid.

func part2(input []string) (output int) {

	maxX, maxY := len(input)-1, len(input[0])-1

	for {

		markedPositions := [][]int{}

		for y, line := range input {
			for x, char := range line {
				if char == '.' {
					continue
				}
				neighbours := getNeighbours(x, y, maxX, maxY)
				if getCount(neighbours, input) < 4 {
					output++
					markedPositions = append(markedPositions, []int{x, y})
				}
			}
		}

		if len(markedPositions) == 0 {
			break
		}

		for _, markedPos := range markedPositions {
			mx, my := markedPos[0], markedPos[1]
			newStr := input[my][:mx] + "." + input[my][mx+1:]
			input[my] = newStr
		}
	}

	return
}

func main() {
	input := getInput()
	fmt.Printf("[Part 1] Ans: %v\n", part1(input))
	fmt.Printf("[Part 2] Ans: %v\n", part2(input))
}
