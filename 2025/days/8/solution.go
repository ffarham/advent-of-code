package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	x, y, z float64
}

func (v *Position) distance(other Position) (output float64) {
	deltaXS := math.Pow(v.x-other.x, 2)
	deltaYS := math.Pow(v.y-other.y, 2)
	deltaZS := math.Pow(v.z-other.z, 2)
	output += deltaXS + deltaYS + deltaZS
	return
}

type Pair struct {
	pos1, pos2 Position
}

func (v *Pair) distance() float64 {
	return v.pos1.distance(v.pos2)
}

func getInput() (output []Position) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		output = append(output, Position{float64(x), float64(y), float64(z)})
	}
	return
}

func part1(input []Position) (output int) {

	// build a slice of all pairs
	pairs := []Pair{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			pair := Pair{input[i], input[j]}
			pairs = append(pairs, pair)
		}
	}

	// sort slice in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance() < pairs[j].distance()
	})

	// build circuits by making N connections between the closest pairs
	N := 10
	circuitNum := 0
	circuits := make(map[Position]int)
	for i := 0; i < N; i++ {
		pair := pairs[i]
		if val1, ok1 := circuits[pair.pos1]; ok1 {
			if val2, ok2 := circuits[pair.pos2]; !ok2 {
				circuits[pair.pos2] = val1
			} else {
				if val1 < val2 {
					for k, v := range circuits {
						if v == val2 {
							circuits[k] = val1
						}
					}
				} else if val1 > val2 {
					for k, v := range circuits {
						if v == val1 {
							circuits[k] = val2
						}
					}
				}
			}
		} else {
			if val2, ok2 := circuits[pair.pos2]; ok2 {
				circuits[pair.pos1] = val2
			} else {
				circuits[pair.pos1] = circuitNum
				circuits[pair.pos2] = circuitNum
				circuitNum++
			}
		}
	}

	// determine the size of each circuit
	circuitCounts := make(map[int]int)
	for _, circuitNum := range circuits {
		if _, ok := circuitCounts[circuitNum]; !ok {
			circuitCounts[circuitNum] = 0
		}
		circuitCounts[circuitNum]++
	}

	type kv struct{ k, v int }
	circuitCountsArr := []kv{}
	for k, v := range circuitCounts {
		circuitCountsArr = append(circuitCountsArr, kv{k, v})
	}
	sort.Slice(circuitCountsArr, func(i, j int) bool {
		return circuitCountsArr[i].v > circuitCountsArr[j].v
	})
	output = 1
	for i := 0; i < 3; i++ {
		output *= circuitCountsArr[i].v
	}

	return
}

func part2(input []Position) (output float64) {
	// build a slice of all pairs
	pairs := []Pair{}
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			pair := Pair{input[i], input[j]}
			pairs = append(pairs, pair)
		}
	}

	// sort slice in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].distance() < pairs[j].distance()
	})

	// determine the last connection
	var lastConn Pair
	circuitNum := 0
	circuits := make(map[Position]int)
	for _, pair := range pairs {
		if val1, ok1 := circuits[pair.pos1]; ok1 {
			if val2, ok2 := circuits[pair.pos2]; !ok2 {
				lastConn = pair
				circuits[pair.pos2] = val1
			} else {
				if val1 < val2 {
					lastConn = pair
					for k, v := range circuits {
						if v == val2 {
							circuits[k] = val1
						}
					}
				} else if val1 > val2 {
					lastConn = pair
					for k, v := range circuits {
						if v == val1 {
							circuits[k] = val2
						}
					}
				}
			}
		} else {
			if val2, ok2 := circuits[pair.pos2]; ok2 {
				lastConn = pair
				circuits[pair.pos1] = val2
			} else {
				circuits[pair.pos1] = circuitNum
				circuits[pair.pos2] = circuitNum
				circuitNum++
			}
		}
	}

	output = lastConn.pos1.x * lastConn.pos2.x
	return
}

func main() {
	input := getInput()
	fmt.Printf("[Part 1] Ans: %v\n", part1(input))
	fmt.Printf("[Part 2] Ans: %v\n", part2(input))
}
