package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseInputPart1(input []string) (output [][]string) {
	for _, line := range input {
		arr := strings.Split(line, " ")
		newArr := []string{}
		for _, elem := range arr {
			if elem == "" {
				continue
			}
			newArr = append(newArr, elem)
		}
		output = append(output, newArr)
	}
	return
}

func part1(input []string) (output int) {
	parsedInput := parseInputPart1(input)
	for i := 0; i < len(parsedInput[0]); i++ {
		isAdd := parsedInput[len(parsedInput)-1][i] == "+"
		runningValue := 0
		if !isAdd {
			runningValue = 1
		}
		for _, row := range parsedInput[:len(parsedInput)-1] {
			numStr := row[i]
			num, _ := strconv.Atoi(numStr)
			if isAdd {
				runningValue += num
			} else {
				runningValue *= num
			}
		}
		output += runningValue
	}
	return output
}

func parseInputPart2(input []string) (output [][]string) {
	newInput := []string{}

	operators := input[len(input)-1]
	operators += " "
	for _, line := range input[:len(input)-1] {
		newLine := ""
		for cIndex, char := range line {
			if char == ' ' && operators[cIndex+1] == ' ' {
				newLine += "."
			} else {
				newLine += string(char)
			}
		}
		newInput = append(newInput, newLine)
	}

	newInput = append(newInput, operators)
	output = parseInputPart1(newInput)
	return
}

func part2(input []string) (output int) {
	parsedInput := parseInputPart2(input)

	for i := 0; i < len(parsedInput[0]); i++ {
		isAdd := parsedInput[len(parsedInput)-1][i] == "+"
		runningValue := 1
		if isAdd {
			runningValue = 0
		}
		for x := 0; x < len(parsedInput[0][i]); x++ {
			numStr := ""
			for y := 0; y < len(parsedInput)-1; y++ {
				if parsedInput[y][i][x] == '.' {
					continue
				}
				numStr += string(parsedInput[y][i][x])
			}
			num, _ := strconv.Atoi(numStr)
			if isAdd {
				runningValue += num
			} else {
				runningValue *= num
			}
		}
		output += runningValue
	}

	return
}

func main() {

	input := getInput()
	fmt.Printf("[Part 1] Ans: %v\n", part1(input))
	fmt.Printf("[Part 2] Ans: %v\n", part2(input))
}
