package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (v Point) directionFrom(other Point) rune {
	if other.x-v.x == 0 {
		if other.y-v.y > 0 {
			return 'N'
		} else {
			return 'S'
		}
	} else if other.y-v.y == 0 {
		if other.x-v.x > 0 {
			return 'W'
		} else {
			return 'E'
		}
	}
	log.Panicln("got the same point")
	return '0'
}

func getInput() (output []Point) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ",")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		output = append(output, Point{x, y})
	}
	return
}

func part1(input []Point) (output int64) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			point1, point2 := input[i], input[j]
			deltaX := point1.x - point2.x + 1
			deltaY := point1.y - point2.y + 1
			area := int64(deltaX) * int64(deltaY)
			if area < 0 {
				area *= -1
			}
			if area > output {
				output = area
			}
		}
	}
	return
}

type Grid struct {
	minX, minY, maxX, maxY int
}

func getGridBounds(points []Point) Grid {
	minX, minY, maxX, maxY := points[0].x, points[0].y, points[0].x, points[0].y
	for _, point := range points[1:] {
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	minX -= 2
	minY -= 2
	maxX += 2
	maxY += 2
	return Grid{minX, minY, maxX, maxY}
}

func getRedTilesMap(points []Point) map[Point]bool {
	redTilesMap := make(map[Point]bool)
	for _, point := range points {
		redTilesMap[point] = true
	}
	return redTilesMap
}

func getGreenTilesMap(points []Point) map[Point]bool {
	greenTilesMap := make(map[Point]bool)

	for i := 1; i < len(points); i++ {
		prevRedTile := points[i-1]
		redTile := points[i]
		if redTile.x-prevRedTile.x == 0 {
			if redTile.y < prevRedTile.y {
				for y := redTile.y + 1; y < prevRedTile.y; y++ {
					p := Point{redTile.x, y}
					greenTilesMap[p] = true
				}
			} else {
				for y := prevRedTile.y + 1; y < redTile.y; y++ {
					p := Point{redTile.x, y}
					greenTilesMap[p] = true
				}
			}
		} else {
			if redTile.x < prevRedTile.x {
				for x := redTile.x + 1; x < prevRedTile.x; x++ {
					p := Point{x, redTile.y}
					greenTilesMap[p] = true
				}
			} else {
				for x := prevRedTile.x + 1; x < redTile.x; x++ {
					p := Point{x, redTile.y}
					greenTilesMap[p] = true
				}
			}
		}
	}
	return greenTilesMap
}

func getStartingPoint(minX, minY, maxX, maxY int, redTiles map[Point]bool) Point {
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			p := Point{x, y}
			if _, ok := redTiles[p]; ok {
				return p
			}
		}
	}
	log.Panicln("failed to find the starting point")
	return Point{}
}

func getFourCorners(point1, point2 Point) []Point {
	if point1.x == point2.x {
		if point1.y < point2.y {
			return []Point{{point1.x, point1.y}, {point1.x, point1.y}, {point1.x, point2.y}, {point1.x, point2.y}}
		} else if point1.y > point2.y {
			return []Point{{point1.x, point2.y}, {point1.x, point2.y}, {point1.x, point1.y}, {point1.x, point1.y}}
		} else {
			log.Panicln("Points 1 and 2 have the same coords", point1, point2)
			return []Point{}
		}
	} else if point1.x < point2.x {
		if point1.y < point2.y {
			return []Point{{point1.x, point1.y}, {point2.x, point1.y}, {point2.x, point2.y}, {point1.x, point2.y}}
		} else if point1.y == point2.y {
			return []Point{{point1.x, point1.y}, {point2.x, point1.y}, {point2.x, point1.y}, {point1.x, point1.y}}
		} else {
			return []Point{{point1.x, point2.y}, {point2.x, point2.y}, {point2.x, point1.y}, {point1.x, point1.y}}
		}
	} else {
		if point1.y < point2.y {
			return []Point{{point2.x, point1.y}, {point1.x, point1.y}, {point1.x, point2.y}, {point2.x, point2.y}}
		} else if point1.y == point2.y {
			return []Point{{point2.x, point1.y}, {point1.x, point1.y}, {point1.x, point1.y}, {point2.x, point1.y}}
		} else {
			return []Point{{point2.x, point2.y}, {point1.x, point2.y}, {point1.x, point1.y}, {point2.x, point1.y}}
		}
	}
}

func part2(points []Point) (output int) {

	grid := getGridBounds(points)
	minX, minY, maxX, maxY := grid.minX, grid.minY, grid.maxX, grid.maxY

	redTiles := getRedTilesMap(points)

	// connect the last point to the first point to determine the green tiles
	points = append(points, points[0])
	greenTiles := getGreenTilesMap(points)
	points = points[:len(points)-1]

	// starting point is the left most point that is the highest in the grid
	startingPoint := getStartingPoint(minX, minY, maxX, maxY, redTiles)

	log.Println("Building a boundry around the red and green tiles")

	// wrap around the starting point to get on the north side
	boundary := make(map[Point]bool)
	writePointer := Point{startingPoint.x - 1, startingPoint.y}
	boundary[writePointer] = true
	writePointer.y--
	boundary[writePointer] = true
	writePointer.x++
	boundary[writePointer] = true

	prevPointer := Point{startingPoint.x, startingPoint.y}
	pointer := Point{startingPoint.x + 1, startingPoint.y}
	writePointer.x++

	// follow the red and green tiles around until the starting point is reached
	for pointer != startingPoint {
		direction := pointer.directionFrom(prevPointer)
		_, isPointerRed := redTiles[pointer]

		if !isPointerRed {
			boundary[writePointer] = true
			if direction == 'E' {
				pointer.x++
				writePointer.x++
			} else if direction == 'S' {
				pointer.y++
				writePointer.y++
			} else if direction == 'W' {
				pointer.x--
				writePointer.x--
			} else {
				pointer.y--
				writePointer.y--
			}
		} else {
			var left, straight, right Point
			if direction == 'E' {
				left = Point{pointer.x, pointer.y - 1}
				straight = Point{pointer.x + 1, pointer.y}
				right = Point{pointer.x, pointer.y + 1}
			} else if direction == 'S' {
				left = Point{pointer.x + 1, pointer.y}
				straight = Point{pointer.x, pointer.y + 1}
				right = Point{pointer.x - 1, pointer.y}
			} else if direction == 'W' {
				left = Point{pointer.x, pointer.y + 1}
				straight = Point{pointer.x - 1, pointer.y}
				right = Point{pointer.x, pointer.y - 1}
			} else if direction == 'N' {
				left = Point{pointer.x - 1, pointer.y}
				straight = Point{pointer.x, pointer.y - 1}
				right = Point{pointer.x + 1, pointer.y}
			}

			_, isLeftRed := redTiles[left]
			_, isLeftGreen := greenTiles[left]
			_, isStraightRed := redTiles[straight]
			_, isStraightGreen := greenTiles[straight]
			_, isRightRed := redTiles[right]
			_, isRightGreen := greenTiles[right]

			if isLeftGreen || isLeftRed {
				if direction == 'E' {
					prevPointer = pointer
					pointer = Point{pointer.x, pointer.y - 1}
					writePointer.x--
				} else if direction == 'S' {
					prevPointer = pointer
					pointer = Point{pointer.x + 1, pointer.y}
					writePointer.y--
				} else if direction == 'W' {
					prevPointer = pointer
					pointer = Point{pointer.x, pointer.y + 1}
					writePointer.x++
				} else {
					prevPointer = pointer
					pointer = Point{pointer.x - 1, pointer.y}
					writePointer.y++
				}
			} else if isStraightGreen || isStraightRed {
				boundary[writePointer] = true
				if direction == 'E' {
					writePointer.x++
					pointer.x++
				} else if direction == 'S' {
					writePointer.y++
					pointer.y++
				} else if direction == 'W' {
					writePointer.x--
					pointer.x--
				} else {
					writePointer.y--
					pointer.y--
				}
			} else if isRightGreen || isRightRed {
				if direction == 'E' {
					boundary[writePointer] = true
					writePointer.x++
					boundary[writePointer] = true
					writePointer.y++
					boundary[writePointer] = true

					prevPointer = pointer
					pointer = Point{prevPointer.x, prevPointer.y + 1}
					writePointer.y++
				} else if direction == 'S' {
					boundary[writePointer] = true
					writePointer.y++
					boundary[writePointer] = true
					writePointer.x--
					boundary[writePointer] = true

					prevPointer = pointer
					pointer = Point{prevPointer.x - 1, prevPointer.y}
					writePointer.x--
				} else if direction == 'W' {
					boundary[writePointer] = true
					writePointer.x--
					boundary[writePointer] = true
					writePointer.y--
					boundary[writePointer] = true

					prevPointer = pointer
					pointer = Point{prevPointer.x, prevPointer.y - 1}
					writePointer.y--
				} else {
					boundary[writePointer] = true
					writePointer.y--
					boundary[writePointer] = true
					writePointer.x++
					boundary[writePointer] = true

					prevPointer = pointer
					pointer = Point{prevPointer.x + 1, prevPointer.y}
					writePointer.x++
				}
			} else {
				log.Panicln("Could not find red or green tiles left, straight, or right of the current point")
			}
		}
	}

	// iterating through all possible rectangles to find one with the largest area
	// such that its perimeter does not cross the invalid boundary
	log.Println("Search for the rectangles...")

	// NOTE: this can potentially be optimised using coordinate compression
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			point1, point2 := points[i], points[j]
			corners := getFourCorners(point1, point2)

			failed := false

			// top side
			for x := corners[0].x; x <= corners[1].x; x++ {
				point := Point{x, corners[0].y}
				if _, ok := boundary[point]; ok {
					failed = true
					break
				}
			}
			if failed {
				continue
			}

			// right side
			for y := corners[1].y; y <= corners[2].y; y++ {
				point := Point{corners[1].x, y}
				if _, ok := boundary[point]; ok {
					failed = true
					break
				}
			}
			if failed {
				continue
			}

			// bottom side
			for x := corners[3].x; x <= corners[2].x; x++ {
				point := Point{x, corners[2].y}
				if _, ok := boundary[point]; ok {
					failed = true
					break
				}
			}
			if failed {
				continue
			}

			// left side
			for y := corners[0].y; y <= corners[3].y; y++ {
				point := Point{corners[0].x, y}
				if _, ok := boundary[point]; ok {
					failed = true
					break
				}
			}
			if failed {
				continue
			}

			width := point1.x - point2.x
			if width < 0 {
				width *= -1
			}
			width++
			height := point1.y - point2.y
			if height < 0 {
				height *= -1
			}
			height++

			area := width * height
			if area > output {
				output = area
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
