package day13

import (
	"errors"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day13/input.txt"
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

// Find x*x + 3*x + 2*x*y + y + y*y.
// Add the office designer's favorite number (your puzzle input).
// Find the binary representation of that sum; count the number of bits that are 1.
// If the number of bits that are 1 is even, it's an open space.
// If the number of bits that are 1 is odd, it's a wall.

func isWall(x int, y int, designNumber int) bool {
	num := bits.OnesCount(uint(x*x + 3*x + 2*x*y + y + y*y + designNumber))
	if num%2 == 0 {
		return false
	} else {
		return true
	}
}

func printGrid(maxX int, maxY int, designNumber int) {
	strBuilder := strings.Builder{}
	for y := range maxY {
		for x := range maxX {
			if isWall(x, y, designNumber) {
				strBuilder.WriteRune('#')
			} else {
				strBuilder.WriteRune('.')
			}
		}
		strBuilder.WriteString("\n")
	}
	fmt.Print(strBuilder.String())
}

func findStepsToTarget(x int, y int, targetX int, targetY int, designNumber int, logger *log.Logger) int {
	queue := [][3]int{{x, y, 0}}
	visited := map[string]bool{}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		currX := queue[0][0]
		currY := queue[0][1]
		currSteps := queue[0][2]
		queue = queue[1:]

		// Add current to visited map
		currKey := fmt.Sprintf("%d,%d", currX, currY)
		visited[currKey] = true

		// Return if at the target
		for currX == targetX && currY == targetY {
			return currSteps
		}

		// Add each direction to the queue if can move to it
		for _, d := range dirs {
			nextX := currX + d[0]
			nextY := currY + d[1]
			nextKey := fmt.Sprintf("%d,%d", nextX, nextY)
			nextSteps := currSteps + 1

			if !isWall(nextX, nextY, designNumber) && !visited[nextKey] && nextX > 0 && nextY > 0 {
				queue = append(queue, [3]int{nextX, nextY, nextSteps})
			}
		}

	}
	return -1 // No path was found
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	designNumber, err := strconv.Atoi(string(f))
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("Office designers favorite number: %d", designNumber)

	var targetX int
	var targetY int
	if strings.HasPrefix(inputFile, "example") {
		targetX = 7
		targetY = 4
	} else {
		targetX = 31
		targetY = 39
	}
	startX := 1
	startY := 1

	printGrid(targetX+3, targetY+3, designNumber)
	logger.Printf("Finding shortest route from (%d, %d) -> (%d, %d)", startX, startY, targetX, targetY)
	answer := findStepsToTarget(startX, startY, targetX, targetY, designNumber, logger)
	return answer
}

func moveSteps(x int, y int, designNumber int, maxSteps int, visited map[string]bool, logger *log.Logger) {
	queue := [][3]int{{x, y, 0}}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		currX := queue[0][0]
		currY := queue[0][1]
		currSteps := queue[0][2]
		queue = queue[1:]

		// Add current to visited map
		currKey := fmt.Sprintf("%d,%d", currX, currY)
		visited[currKey] = true

		// Return if at max steps
		if currSteps >= maxSteps {
			continue
		}

		// Add each direction to the queue if can move to it
		for _, d := range dirs {
			nextX := currX + d[0]
			nextY := currY + d[1]
			nextSteps := currSteps + 1
			nextKey := fmt.Sprintf("%d,%d", nextX, nextY)

			if !isWall(nextX, nextY, designNumber) && !visited[nextKey] && nextX >= 0 && nextY >= 0 {
				queue = append(queue, [3]int{nextX, nextY, nextSteps})
			}
		}

	}
	return
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	designNumber, err := strconv.Atoi(string(f))
	if err != nil {
		logger.Fatal(err)
	}
	logger.Printf("Office designers favorite number: %d", designNumber)

	startX := 1
	startY := 1
	maxSteps := 50

	printGrid(10, 10, designNumber)
	logger.Printf("Finding how many locations we can visit in %d steps", maxSteps)
	visited := map[string]bool{}
	moveSteps(startX, startY, designNumber, maxSteps, visited, logger)
	answer := len(visited)
	return answer
}
