package day24

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day24/input.txt"
	if len(inputFile) > 0 {
		file = inputFile[0]
	}

	switch part {
	case 1:
		return SolvePart1(file, logger), nil
	case 2:
		return SolvePart2(file, logger), nil
	default:
		return -1, errors.New("incorrect part number recieved")
	}
}

func printGrid(grid [][]rune, logger *log.Logger) {
	str := "\n"
	for _, row := range grid {
		str += fmt.Sprintf("%c\n", row)
	}
	logger.Print(str)
}

type State struct {
	x, y         int
	dist         int
	visited_mask int
}

type VisitedKey struct {
	x, y         int
	visited_mask int
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	grid := [][]rune{}
	points := map[int][2]int{}
	scanner := bufio.NewScanner(f)
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()
		row := []rune{}
		for x, c := range line {
			row = append(row, c)
			if c >= '0' && c <= '9' {
				num, _ := strconv.Atoi(string(c))
				points[num] = [2]int{x, y}
			}
		}
		grid = append(grid, row)
	}

	// Print for debugging
	printGrid(grid, logger)
	logger.Print("Points of interest")
	for _, c := range []int{0, 1, 2, 3, 4} {
		logger.Printf("%d : (%d, %d)", c, points[c][0], points[c][1])
	}

	answer := 0

	// Initialise BFS search
	queue := []State{{x: points[0][0], y: points[0][1], dist: 0, visited_mask: 1}}
	visited := map[VisitedKey]bool{{x: points[0][0], y: points[0][1], visited_mask: 1}: true}
	targetMask := (1 << len(points)) - 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// logger.Printf("x: %d,y: %d, dist: %d, mask: %b", curr.x, curr.y, curr.dist, curr.visited_mask)

		// Check if end state is reached
		if curr.visited_mask == targetMask {
			answer = curr.dist
			break
		}

		dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range dirs {
			nextX, nextY := curr.x+d[0], curr.y+d[1]
			if grid[nextY][nextX] == '#' {
				continue
			}

			nextMask := curr.visited_mask
			for p, pos := range points {
				if pos[0] == nextX && pos[1] == nextY {
					nextMask |= 1 << p
					break
				}
			}

			key := VisitedKey{nextX, nextY, nextMask}
			if !visited[key] {
				visited[key] = true
				queue = append(queue, State{x: nextX, y: nextY, dist: curr.dist + 1, visited_mask: nextMask})
			}

		}
	}

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	grid := [][]rune{}
	points := map[int][2]int{}
	scanner := bufio.NewScanner(f)
	y := -1
	for scanner.Scan() {
		y++
		line := scanner.Text()
		row := []rune{}
		for x, c := range line {
			row = append(row, c)
			if c >= '0' && c <= '9' {
				num, _ := strconv.Atoi(string(c))
				points[num] = [2]int{x, y}
			}
		}
		grid = append(grid, row)
	}

	// Print for debugging
	printGrid(grid, logger)
	logger.Print("Points of interest")
	for _, c := range []int{0, 1, 2, 3, 4} {
		logger.Printf("%d : (%d, %d)", c, points[c][0], points[c][1])
	}

	answer := 0

	// Initialise BFS search
	queue := []State{{x: points[0][0], y: points[0][1], dist: 0, visited_mask: 1}}
	visited := map[VisitedKey]bool{{x: points[0][0], y: points[0][1], visited_mask: 1}: true}
	targetMask := (1 << len(points)) - 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// logger.Printf("x: %d,y: %d, dist: %d, mask: %b", curr.x, curr.y, curr.dist, curr.visited_mask)

		// Check if visited all points and ended at start
		if curr.visited_mask == targetMask && curr.x == points[0][0] && curr.y == points[0][1] {
			answer = curr.dist
			break
		}

		dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range dirs {
			nextX, nextY := curr.x+d[0], curr.y+d[1]
			if grid[nextY][nextX] == '#' {
				continue
			}

			nextMask := curr.visited_mask
			for p, pos := range points {
				if pos[0] == nextX && pos[1] == nextY {
					nextMask |= 1 << p
					break
				}
			}

			key := VisitedKey{nextX, nextY, nextMask}
			if !visited[key] {
				visited[key] = true
				queue = append(queue, State{x: nextX, y: nextY, dist: curr.dist + 1, visited_mask: nextMask})
			}

		}
	}

	return answer
}
