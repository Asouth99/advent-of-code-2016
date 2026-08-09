package day06

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day06/input.txt"
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

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	columns := []string{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i == 0 {
			for range len(line) {
				columns = append(columns, "")
			}
		}
		for idx, char := range line {
			columns[idx] += string(char)
		}
	}

	answer := ""
	for _, str := range columns {
		// Get char with most occurrences in str
		charsMap := map[string]int{}
		for _, char := range str {
			charsMap[string(char)]++
		}
		logger.Printf("Chars: %+v", charsMap)
		max := 0
		maxKey := ""
		for key, val := range charsMap {
			if val > max {
				max = val
				maxKey = key
			}
		}
		answer += maxKey
	}

	logger.Print(columns)
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	columns := []string{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i == 0 {
			for range len(line) {
				columns = append(columns, "")
			}
		}
		for idx, char := range line {
			columns[idx] += string(char)
		}
	}

	answer := ""
	for _, str := range columns {
		// Get char with most occurrences in str
		charsMap := map[string]int{}
		for _, char := range str {
			charsMap[string(char)]++
		}
		logger.Printf("Chars: %+v", charsMap)
		min := math.MaxInt64
		minKey := ""
		for key, val := range charsMap {
			if val < min {
				min = val
				minKey = key
			}
		}
		answer += minKey
	}

	logger.Print(columns)
	return answer
}
