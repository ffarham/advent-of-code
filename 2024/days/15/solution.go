package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WALL           = '#'
	FREE           = '.'
	OBSTACLE       = 'O'
	PLAYER_CHAR    = '@'
	OPEN_OBSTACLE  = '['
	CLOSE_OBSTACLE = ']'
)

type location struct {
	x, y int
}

type player struct {
	loc location
}

func (v *player) nextLocation(move rune) location {
	if move == '^' {
		return location{v.loc.x, v.loc.y - 1}
	} else if move == '>' {
		return location{v.loc.x + 1, v.loc.y}
	} else if move == 'v' {
		return location{v.loc.x, v.loc.y + 1}
	} else if move == '<' {
		return location{v.loc.x - 1, v.loc.y}
	}
	panic("invalid move")
}

func (v *player) step(loc location) {
	v.loc = loc
}

type world [][]rune

func (v *world) isWall(loc location) bool {
	return WALL == (*v)[loc.y][loc.x]
}

func (v *world) isFree(loc location) bool {
	return FREE == (*v)[loc.y][loc.x]
}

func (v *world) isObstaclePart1(loc location) bool {
	return OBSTACLE == (*v)[loc.y][loc.x]
}

func (v *world) isObstaclepart2(loc location) bool {
	return OPEN_OBSTACLE == (*v)[loc.y][loc.x] || CLOSE_OBSTACLE == (*v)[loc.y][loc.x]
}

func (v *world) getFurthestObstaclePart1(loc, pLoc location) (location, bool) {
	xDelta, yDelta := loc.x-pLoc.x, loc.y-pLoc.y
	ptr := loc
	for v.isObstaclePart1(ptr) {
		ptr = location{ptr.x + xDelta, ptr.y + yDelta}
	}
	atWall := WALL == (*v)[ptr.y][ptr.x]
	return location{ptr.x - xDelta, ptr.y - yDelta}, atWall
}

func (v *world) getFurthestObstaclePart2(loc, pLoc location) (location, bool) {
	xDelta, yDelta := loc.x-pLoc.x, loc.y-pLoc.y

	if OPEN_OBSTACLE == (*v)[loc.y][loc.x] {
		getFurthestObstaclePart2Helper(v, loc, location{loc.x + 1, loc.y}, xDelta, yDelta)
	} else if CLOSE_OBSTACLE == (*v)[loc.y][loc.x] {
		getFurthestObstaclePart2Helper(v, loc, location{loc.x - 1, loc.y}, xDelta, yDelta)
	}

	panic("invalid obstacle when determining between open and close")
}

func getFurthestObstaclePart2Helper(grid *world, open, close location, xDelta, yDelta int) (location, location, bool, bool) {

	// check if movement is horizontal or vertical
	if xDelta != 0 && yDelta != 0 {
		panic("player can not move diagonally")
	} else if xDelta != 0 {
		// horixontal movement

		if OPEN_OBSTACLE != (*grid)[open.y][open.x] || CLOSE_OBSTACLE != (*grid)[close.y][close.x] {

			var isWall bool
			if xDelta > 0 {
				// moving right, so check the left cell for wall
				isWall = WALL == (*grid)[open.y][open.x]
			} else {
				// moving left, so check the right cell for wall
				isWall = WALL == (*grid)[close.y][close.x]
			}
			bothFree := FREE == (*grid)[open.y][open.x] && FREE == (*grid)[close.y][close.x]
			return location{open.x - 2*xDelta, open.y - 2*yDelta}, location{close.x - 2*xDelta, close.y - 2*yDelta}, isWall, bothFree
		}
		return getFurthestObstaclePart2Helper(grid, location{open.x + 2*xDelta, open.y + 2*yDelta}, location{close.x + 2*xDelta, close.y + 2*yDelta})

	} else if yDelta != 0 {
		// vertical movement
		if OPEN_OBSTACLE == (*grid)[open.y][open.x] && CLOSE_OBSTACLE == (*grid)[close.y][close.x] {
			newOpen, newClose := location{open.x, open.y + yDelta}, location{close.x, close.y + yDelta}
			// case: 	[]
			// 			[]
			if OPEN_OBSTACLE == (*grid)[newOpen.y][newOpen.x] && CLOSE_OBSTACLE == (*grid)[newClose.y][newClose.x] {
				return getFurthestObstaclePart2Helper(grid, newOpen, newClose, xDelta, yDelta)
			} else if OPEN_OBSTACLE == (*grid)[newOpen.y][newOpen.x+1] && CLOSE_OBSTACLE == (*grid)[newClose.y][newClose.x+1] && OPEN_OBSTACLE == (*grid)[newOpen.y][newOpen.x-1] && CLOSE_OBSTACLE == (*grid)[newClose.y][newClose.x-1] {
				// case:   [][]		[]
				// 			[]	   [][]
				// TODO: return a list of furthest points instead
				fOpenR, fCloseR, isWallR, bothFreeR := getFurthestObstaclePart2Helper(grid, location{newOpen.x + 1, newOpen.y}, location{newClose.x + 1, newClose.y}, xDelta, yDelta)
				fOpenL, fCloseL, isWallL, bothFreeL := getFurthestObstaclePart2Helper(grid, location{newOpen.x - 1, newOpen.y}, location{newClose.x - 1, newClose.y}, xDelta, yDelta)

			} else if OPEN_OBSTACLE == (*grid)[newOpen.y][newOpen.x+1] && CLOSE_OBSTACLE == (*grid)[newClose.y][newClose.x+1] {
				// case:     []		[]
				// 			[]		 []
				return getFurthestObstaclePart2Helper(grid, location{newOpen.x + 1, newOpen.y}, location{newClose.x + 1, newClose.y}, xDelta, yDelta)
			} else if OPEN_OBSTACLE == (*grid)[newOpen.y][newOpen.x-1] && CLOSE_OBSTACLE == (*grid)[newClose.y][newClose.x-1] {
				// case:   	[]	  	[]
				// 			 []	   []
				return getFurthestObstaclePart2Helper(grid, location{newOpen.x - 1, newOpen.y}, location{newClose.x - 1, newClose.y}, xDelta, yDelta)
			} else {
				// check if wall: base case
			}

		} else {
			panic("expecting to be at obstacle")
		}
	}
	panic("player should be moving somewhere")
}

func (v *world) moveObstaclesPart1(fLoc, loc, pLoc location) {
	xDelta, yDelta := loc.x-pLoc.x, loc.y-pLoc.y

	newLoc := location{fLoc.x + xDelta, fLoc.y + yDelta}
	if FREE != (*v)[newLoc.y][newLoc.x] {
		panic("the location the obstacles are being shifted to is not free")
	}

	(*v)[newLoc.y][newLoc.x] = OBSTACLE
	(*v)[loc.y][loc.x] = FREE
}

func getContents(double bool) (world, []rune) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	moves := []rune{}
	grid := world{}

	gridFlag := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			gridFlag = false
			continue
		}

		if gridFlag {
			row := []rune{}
			for _, char := range line {
				if double {
					if WALL == char {
						row = append(row, WALL)
						row = append(row, WALL)
					} else if OBSTACLE == char {
						row = append(row, OPEN_OBSTACLE)
						row = append(row, CLOSE_OBSTACLE)
					} else if FREE == char {
						row = append(row, FREE)
						row = append(row, FREE)
					} else if PLAYER_CHAR == char {
						row = append(row, PLAYER_CHAR)
						row = append(row, FREE)
					} else {
						panic("picked up invalid char in double")
					}
				} else {
					row = append(row, char)
				}
			}
			grid = append(grid, row)
		} else {
			moves = append(moves, []rune(line)...)
		}
	}
	return grid, moves
}

func getPLayer(grid *world) player {
	for y := 0; y < len((*grid)); y++ {
		for x := 0; x < len((*grid)[0]); x++ {
			if (*grid)[y][x] == PLAYER_CHAR {
				(*grid)[y][x] = FREE
				return player{location{x, y}}
			}
		}
	}
	panic("cant find player on grid")
}

func simulatePart1(grid *world, moves []rune) {
	p := getPLayer(grid)

	for _, move := range moves {

		// plot(grid, p)

		newLoc := p.nextLocation(move)

		// if would collide with wall, dont move player
		if grid.isWall(newLoc) {
			continue
		}

		// step forward if free
		if grid.isFree(newLoc) {
			p.step(newLoc)
			continue
		}

		// handle obstacles
		if grid.isObstaclePart1(newLoc) {
			foLoc, atWall := grid.getFurthestObstaclePart1(newLoc, p.loc)
			if atWall {
				// if chain of obstacles hits a wall, dont move player
				continue
			}

			// shift obstacles
			grid.moveObstaclesPart1(foLoc, newLoc, p.loc)
			p.step(newLoc)
			continue
		}

		panic("unidentified object at new location")

	}
}

func addCoords(grid *world) int {

	output := 0

	for y := 0; y < len((*grid)); y++ {
		for x := 0; x < len((*grid)[0]); x++ {
			if OBSTACLE == (*grid)[y][x] {
				output += (100 * y) + x
			}
		}
	}
	return output
}

func plot(grid *world, p player) {

	output := ""
	for y, row := range *grid {
		for x, char := range row {
			if x == p.loc.x && y == p.loc.y {
				output += string(PLAYER_CHAR)
			} else {
				output += string(char)
			}
		}
		output += "\n"
	}
	fmt.Println(output)
}

func part1() {
	grid, moves := getContents(false)

	simulatePart1(&grid, moves)
	ans := addCoords(&grid)

	fmt.Println("Part 1 Ans: ", ans)
}

func simulatePart2(grid *world, moves []rune) {

	p := getPLayer(grid)

	for _, move := range moves {
		newLoc := p.nextLocation(move)

		// do nothing if collides with wall
		if grid.isWall(newLoc) {
			continue
		}

		// step forward if free
		if grid.isFree(newLoc) {
			p.step(newLoc)
			continue
		}

		if grid.isObstaclepart2(newLoc) {

		}

	}

}

func part2() {

	grid, moves := getContents(true)

	simulatePart2(&grid, moves)

}

func main() {
	// part1()
	part2()
}
