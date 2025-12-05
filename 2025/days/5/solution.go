package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() (output [][]string) {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fresh := []string{}
	avail := []string{}

	emptyLineFlag := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptyLineFlag = true
			continue
		}
		if !emptyLineFlag {
			fresh = append(fresh, line)
		} else {
			avail = append(avail, line)
		}
	}

	output = append(output, fresh)
	output = append(output, avail)
	return
}

// Time complexity: O(m + nm) where m is the number of fresh ranges and n is the number of
// available items.

func part1(input [][]string) (output int) {
	fresh, avail := input[0], input[1]

	freshRanges := [][]int{}
	for _, freshRange := range fresh {
		lbUbStr := strings.Split(freshRange, "-")
		lbStr, ubStr := lbUbStr[0], lbUbStr[1]
		lb, _ := strconv.Atoi(lbStr)
		ub, _ := strconv.Atoi(ubStr)
		freshRanges = append(freshRanges, []int{lb, ub})
	}

	for _, item := range avail {
		itemNum, _ := strconv.Atoi(item)
		for _, freshRange := range freshRanges {
			if itemNum >= freshRange[0] && itemNum <= freshRange[1] {
				output++
				break
			}
		}
	}
	return
}

// Time complexity: O(m + m^4) where m is the initial number of fresh ranges.
// A range can be split m times in the worst case, so you really have to check m^2 ranges for
// overlaps. The visited list will contain m^2 ranges in the worst case, thus m^2 * m^2.

func part2(input []string) (output int) {

	freshRanges := list.New()
	for _, freshRange := range input {
		lbUbStr := strings.Split(freshRange, "-")
		lbStr, ubStr := lbUbStr[0], lbUbStr[1]
		lb, _ := strconv.Atoi(lbStr)
		ub, _ := strconv.Atoi(ubStr)
		freshRanges.PushBack([]int{lb, ub})
	}

	visited := [][]int{}
	for {
		freshRange := freshRanges.Front()
		if freshRange == nil {
			break
		}
		freshRanges.Remove(freshRange)

		lb, ub := freshRange.Value.([]int)[0], freshRange.Value.([]int)[1]

		splitFlag := false
		for _, visit := range visited {
			vLb, vUb := visit[0], visit[1]

			// range is entirely on the left of the visited range
			if lb < vLb && ub < vLb {
				continue
			}

			// range starts before the visited range
			if lb < vLb {
				splitFlag = true

				// range ends in the visited range
				if ub >= vLb && ub <= vUb {
					freshRanges.PushBack([]int{lb, vLb - 1})
					break
				}
				// range ends after the visited range
				if ub > vUb {
					freshRanges.PushBack([]int{lb, vLb - 1})
					freshRanges.PushBack([]int{vUb + 1, ub})
					break
				}
			}

			// range starts and ends on the visited LB
			if lb == vLb && ub == vLb {
				splitFlag = true // implies delete
				break
			}

			// range starts and ends in the visited range
			if lb >= vLb && ub <= vUb {
				splitFlag = true // implies delete
				break
			}

			// range starts and ends on the visited UB
			if lb == vUb && ub == vUb {
				splitFlag = true // implies delete
				break
			}

			// range starts in the visited range and ends afterwards
			if lb >= vLb && lb <= vUb {
				splitFlag = true
				freshRanges.PushBack([]int{vUb + 1, ub})
				break
			}

			// range is entirely on the right of the visited range
		}

		if !splitFlag {
			visited = append(visited, []int{lb, ub})
		}

	}

	for _, freshRange := range visited {
		output += (freshRange[1] - freshRange[0] + 1)
	}

	return
}

type Point struct {
	value, isStart int
}

type PointSlice []Point

func (v PointSlice) Len() int {
	return len(v)
}

func (v PointSlice) Less(i, j int) bool {
	if v[i].value != v[j].value {
		return v[i].value < v[j].value
	}
	return v[i].isStart == 1
}

func (v PointSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Time complexity: O(mlog(m)) where m is the number of fresh ranges.
// Uses the Sweep line algorithm.

func part2Version2(input []string) (output int) {

	freshRanges := PointSlice{}
	for _, freshRange := range input {
		lbUbStr := strings.Split(freshRange, "-")
		lbStr, ubStr := lbUbStr[0], lbUbStr[1]
		lb, _ := strconv.Atoi(lbStr)
		ub, _ := strconv.Atoi(ubStr)
		freshRanges = append(freshRanges, Point{lb, 1})
		freshRanges = append(freshRanges, Point{ub, -1})
	}

	sort.Sort(freshRanges)

	flag := 0
	anchor := -1
	for _, freshRange := range freshRanges {
		if anchor == -1 {
			anchor = freshRange.value
		}

		flag += freshRange.isStart

		if flag == 0 {
			output += (freshRange.value - anchor + 1)
			anchor = -1
		}
	}
	return
}

func main() {
	input := getInput()
	fmt.Printf("[PART 1] Ans: %v\n", part1(input))
	fmt.Printf("[PART 2] Ans: %v\n", part2(input[0]))
	fmt.Printf("[PART 2 v2] Ans: %v\n", part2Version2(input[0]))
}
