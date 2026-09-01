package day23

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day23/input.txt"
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

func printState(reg map[string]int, ptr int, programLines [][]string, logger *log.Logger) {
	str := "|"
	for _, r := range []string{"a", "b", "c", "d"} {
		str += fmt.Sprintf(" %s:%d |", r, reg[r])
	}
	str += fmt.Sprintf("\nptr: %d - %v\n", ptr, programLines[ptr])
	logger.Print(str)
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
	registers := map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	programPtr := 0

	// Loop over program
	for programPtr < len(programLines) {
		programLine := programLines[programPtr]

		printState(registers, programPtr, programLines, logger)

		op := programLine[0]

		switch op {
		case "tgl":
			n := registers[programLine[1]] + programPtr
			if n >= 0 && n < len(programLines) {
				inst := programLines[n]
				if len(inst) <= 2 {
					if inst[0] == "inc" {
						programLines[n][0] = "dec"
					} else {
						programLines[n][0] = "inc"
					}
				} else {
					if inst[0] == "jnz" {
						programLines[n][0] = "cpy"
					} else {
						programLines[n][0] = "jnz"
					}
				}
			}
		case "cpy":
			if _, ok := registers[programLine[2]]; ok {
				if _, ok := registers[programLine[1]]; ok {
					registers[programLine[2]] = registers[programLine[1]]
				} else {
					val, err := strconv.Atoi(programLine[1])
					if err != nil {
						logger.Fatal(err)
					}
					registers[programLine[2]] = val
				}
			}
		case "inc":
			registers[programLine[1]]++
		case "dec":
			registers[programLine[1]]--
		case "jnz":
			val, ok := registers[programLine[1]]
			if !ok {
				val, err = strconv.Atoi(programLine[1])
				if err != nil {
					logger.Fatal(err)
				}
			}
			if val != 0 {
				jump, ok := registers[programLine[2]]
				if !ok {
					jump, err = strconv.Atoi(programLine[2])
					if err != nil {
						logger.Fatal(err)
					}
				}
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
	registers := map[string]int{"a": 12, "b": 0, "c": 0, "d": 0}
	programPtr := 0

	// Loop over program
	for programPtr < len(programLines) {
		programLine := programLines[programPtr]

		printState(registers, programPtr, programLines, logger)

		op := programLine[0]

		switch op {
		case "tgl":
			n := registers[programLine[1]] + programPtr
			if n >= 0 && n < len(programLines) {
				inst := programLines[n]
				if len(inst) <= 2 {
					if inst[0] == "inc" {
						programLines[n][0] = "dec"
					} else {
						programLines[n][0] = "inc"
					}
				} else {
					if inst[0] == "jnz" {
						programLines[n][0] = "cpy"
					} else {
						programLines[n][0] = "jnz"
					}
				}
			}
		case "cpy":
			if _, ok := registers[programLine[2]]; ok {
				if _, ok := registers[programLine[1]]; ok {
					registers[programLine[2]] = registers[programLine[1]]
				} else {
					val, err := strconv.Atoi(programLine[1])
					if err != nil {
						logger.Fatal(err)
					}
					registers[programLine[2]] = val
				}
			}
		case "inc":
			registers[programLine[1]]++
		case "dec":
			registers[programLine[1]]--
		case "jnz":
			val, ok := registers[programLine[1]]
			if !ok {
				val, err = strconv.Atoi(programLine[1])
				if err != nil {
					logger.Fatal(err)
				}
			}
			if val != 0 {
				jump, ok := registers[programLine[2]]
				if !ok {
					jump, err = strconv.Atoi(programLine[2])
					if err != nil {
						logger.Fatal(err)
					}
				}
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
