package day03

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day03/input.txt"
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

func triangleIsPossible(sides []int) bool {
	// get sum of all sides
	sum := 0
	for _, side := range sides {
		sum += side
	}

	// Check if any side is larger than sum of other sides
	for _, side := range sides {
		sumOfRest := sum - side
		if side >= sumOfRest {
			return false
		}
	}
	return true
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()

		sidesStr := strings.Fields(line)
		sides := []int{}
		for _, sideStr := range sidesStr {
			side, err := strconv.Atoi(sideStr)
			if err != nil {
				logger.Fatal(err)
			}
			sides = append(sides, side)
		}
		logger.Printf("Checking Triangle %d : %v", i, sides)

		isPossible := triangleIsPossible(sides)
		if isPossible {
			answer++
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

	answer := 0
	triangles := [][]int{}

	scanner := bufio.NewScanner(f)
	i := -1
	inputTriangles := [][]int{}
	for scanner.Scan() {
		i++
		line := scanner.Text()

		sidesStr := strings.Fields(line)
		sides := []int{}
		for _, sideStr := range sidesStr {
			side, err := strconv.Atoi(sideStr)
			if err != nil {
				logger.Fatal(err)
			}
			sides = append(sides, side)
		}

		switch i % 3 {
		case 0: // Start of a new triangle
			inputTriangles = [][]int{}
			for _, side := range sides {
				inputTriangles = append(inputTriangles, []int{side})
			}
		case 1:
			for idx, side := range sides {
				inputTriangles[idx] = append(inputTriangles[idx], side)
			}
		case 2: // End of a triangle
			for idx, side := range sides {
				inputTriangles[idx] = append(inputTriangles[idx], side)
			}
			triangles = append(triangles, inputTriangles...)
		}

	}

	logger.Printf("Read %d triangles from input", len(triangles))

	for i, triangle := range triangles {
		logger.Printf("Checking Triangle %d : %v", i, triangle)
		isPossible := triangleIsPossible(triangle)
		if isPossible {
			answer++
		}
	}
	return answer
}
