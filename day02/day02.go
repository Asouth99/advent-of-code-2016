package day02

import (
	internal "aoc2016/internal/utils"
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day02/input.txt"
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

type coord struct {
	x int
	y int
}

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	code := []int{}
	pos := coord{x: 1, y: 1}
	// Bathroom keypad
	//     0 1 2
	//    ------
	// 0 | 1 2 3
	// 1 | 4 5 6
	// 2 | 7 8 9
	keypad := map[int]map[int]int{
		0: {0: 1, 1: 4, 2: 7},
		1: {0: 2, 1: 5, 2: 8},
		2: {0: 3, 1: 6, 2: 9},
	}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		instructions := scanner.Text()
		logger.Printf("Processing instruction %d : %s", i+1, instructions)

		for idx, char := range instructions {
			switch char {
			case 'U':
				pos.y = internal.Max(pos.y-1, 0)
			case 'D':
				pos.y = internal.Min(pos.y+1, 2)
			case 'L':
				pos.x = internal.Max(pos.x-1, 0)
			case 'R':
				pos.x = internal.Min(pos.x+1, 2)
			default:
				logger.Fatalf("Invalid instruction found <%v> at pos %d", char, idx)
			}
		}
		logger.Printf("Digit %d is at position (%d, %d)", i+1, pos.x, pos.y)
		code = append(code, keypad[pos.x][pos.y])

	}
	logger.Print(code)
	codeStr := ""
	for _, digit := range code {
		codeStr += strconv.Itoa(digit)
	}
	answer := codeStr
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	code := ""
	pos := coord{x: 0, y: 2}
	// Bathroom keypad
	//     0 1 2 3 4
	//    -----------
	// 0 |     1
	// 1 |   2 3 4
	// 2 | 5 6 7 8 9
	// 3 |   A B C
	// 4 |     D
	keypad := map[int]map[int]string{
		0: {0: "", 1: "", 2: "1", 3: "", 4: ""},
		1: {0: "", 1: "2", 2: "3", 3: "4", 4: ""},
		2: {0: "5", 1: "6", 2: "7", 3: "8", 4: "9"},
		3: {0: "", 1: "A", 2: "B", 3: "C", 4: ""},
		4: {0: "", 1: "", 2: "D", 3: "", 4: ""},
	}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		instructions := scanner.Text()
		logger.Printf("Processing instruction %d : %s", i+1, instructions)

		for idx, char := range instructions {
			switch char {
			case 'U':
				pos.y = internal.Max(pos.y-1, 0)
			case 'D':
				pos.y = internal.Min(pos.y+1, len(keypad)-1)
			case 'L':
				pos.x = internal.Max(pos.x-1, 0)
			case 'R':
				pos.x = internal.Min(pos.x+1, len(keypad)-1)
			default:
				logger.Fatalf("Invalid instruction found <%v> at pos %d", char, idx)
			}
			// REVERT movement if we're at a position with no number on the keypad
			if keypad[pos.y][pos.x] == "" {
				switch char {
				case 'U':
					pos.y++
				case 'D':
					pos.y--
				case 'L':
					pos.x++
				case 'R':
					pos.x--
				}
			}
		}
		logger.Printf("Digit %d is at position (%d, %d)", i+1, pos.x, pos.y)
		code += keypad[pos.y][pos.x]

	}
	logger.Print(code)
	answer := code
	return answer
}
