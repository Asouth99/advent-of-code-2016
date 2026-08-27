package day18

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day18/input.txt"
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

func printTiles(rows []string) {
	for _, row := range rows {
		fmt.Printf("%s\n", row)
	}
}

// Its left and center tiles are traps, but its right tile is not.
// Its center and right tiles are traps, but its left tile is not.
// Only its left tile is a trap.
// Only its right tile is a trap.
func isTrap(row string, pos int) bool {
	var l, c, r byte
	if pos > 0 {
		l = row[pos-1]
	} else {
		l = '.'
	}
	c = row[pos]
	if pos < len(row)-1 {
		r = row[pos+1]
	} else {
		r = '.'
	}

	if l == '^' && c == '^' && r != '^' {
		return true
	}
	if l != '^' && c == '^' && r == '^' {
		return true
	}
	if l == '^' && c != '^' && r != '^' {
		return true
	}
	if l != '^' && c != '^' && r == '^' {
		return true
	}
	return false
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	var maxRows int
	switch inputFile {
	case "example_1.txt":
		maxRows = 3
	case "example_2.txt":
		maxRows = 10
	default:
		maxRows = 40
	}

	tiles := []string{string(f)}

	logger.Printf("Generating %d rows of tiles", maxRows-1)
	strBuilder := strings.Builder{}
	for i := 1; i < maxRows; i++ {
		strBuilder.Reset()
		strBuilder.Grow(len(tiles[0]))
		for j := range len(tiles[0]) {
			if isTrap(tiles[i-1], j) {
				strBuilder.WriteRune('^')
			} else {
				strBuilder.WriteRune('.')
			}
		}
		tiles = append(tiles, strBuilder.String())
	}

	printTiles(tiles)

	answer := 0
	for _, row := range tiles {
		for _, char := range row {
			if char == '.' {
				answer++
			}
		}
	}

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	var maxRows int
	switch inputFile {
	case "example_1.txt":
		maxRows = 3
	case "example_2.txt":
		maxRows = 10
	default:
		maxRows = 400000
	}

	tiles := []string{string(f)}

	logger.Printf("Generating %d rows of tiles", maxRows-1)
	strBuilder := strings.Builder{}
	for i := 1; i < maxRows; i++ {
		strBuilder.Reset()
		strBuilder.Grow(len(tiles[0]))
		for j := range len(tiles[0]) {
			if isTrap(tiles[i-1], j) {
				strBuilder.WriteRune('^')
			} else {
				strBuilder.WriteRune('.')
			}
		}
		tiles = append(tiles, strBuilder.String())
	}

	// printTiles(tiles)

	answer := 0
	for _, row := range tiles {
		for _, char := range row {
			if char == '.' {
				answer++
			}
		}
	}

	return answer
}
