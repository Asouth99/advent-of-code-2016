package day12

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day12/input.txt"
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

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read program
	programLines := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		programLines = append(programLines, lineSplit)
	}

	// Initialise machine state
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	programPtr := 0

	// Loop over program
	for programPtr < len(programLines) {
		programLine := programLines[programPtr]
		op := programLine[0]

		switch op {
		case "cpy":
			if _, ok := registers[programLine[1]]; ok {
				registers[programLine[2]] = registers[programLine[1]]
			} else {
				val, _ := strconv.Atoi(programLine[1])
				registers[programLine[2]] = val
			}
		case "inc":
			registers[programLine[1]]++
		case "dec":
			registers[programLine[1]]--
		case "jnz":
			val, ok := registers[programLine[1]]
			if !ok {
				val, _ = strconv.Atoi(programLine[1])
			}
			if val != 0 {
				jump, _ := strconv.Atoi(programLine[2])
				programPtr += jump
				continue
			}
		default:
			logger.Fatalf("Unknown operand found in instruction '%s'", programLine)
		}

		programPtr++
	}

	answer := registers["a"]
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read program
	programLines := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		programLines = append(programLines, lineSplit)
	}

	// Initialise machine state
	registers := map[string]int{"a": 0, "b": 0, "c": 1, "d": 0}
	programPtr := 0

	// Loop over program
	for programPtr < len(programLines) {
		programLine := programLines[programPtr]
		op := programLine[0]

		switch op {
		case "cpy":
			if _, ok := registers[programLine[1]]; ok {
				registers[programLine[2]] = registers[programLine[1]]
			} else {
				val, _ := strconv.Atoi(programLine[1])
				registers[programLine[2]] = val
			}
		case "inc":
			registers[programLine[1]]++
		case "dec":
			registers[programLine[1]]--
		case "jnz":
			val, ok := registers[programLine[1]]
			if !ok {
				val, _ = strconv.Atoi(programLine[1])
			}
			if val != 0 {
				jump, _ := strconv.Atoi(programLine[2])
				programPtr += jump
				continue
			}
		default:
			logger.Fatalf("Unknown operand found in instruction '%s'", programLine)
		}

		programPtr++
	}

	answer := registers["a"]
	return answer
}
