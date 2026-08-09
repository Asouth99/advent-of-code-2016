package day01

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day01/input.txt"
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
	type player struct {
		direction int // 0,1,2,3 = N,E,S,W
		x         int
		y         int
	}

	// Read input file
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	// Read input into list of instructions
	var instructions []string
	instructions = strings.Split(string(f), ", ")
	logger.Print(instructions)

	// Initialize starting pos
	me := player{direction: 0, x: 0, y: 0}
	logger.Printf("Me: %+v", me)

	// Process each instruction
	for i, inst := range instructions {
		logger.Printf("%d: Moving %s", i, inst)

		// Rotate ME
		rotation := string(inst[0])
		switch rotation {
		case "R":
			me.direction = (me.direction + 1) % 4
		case "L":
			me.direction = ((me.direction-1)%4 + 4) % 4
		default:
			logger.Fatalf("Unknown rotation found in instruction: %v", inst)
		}

		// Move ME
		numStr := inst[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			logger.Fatal(err)
		}
		switch me.direction {
		case 0:
			me.y = me.y + num
		case 1:
			me.x = me.x + num
		case 2:
			me.y = me.y - num
		case 3:
			me.x = me.x - num
		default:
			logger.Fatalf("Unknown player direction: %d", me.direction)
		}
		logger.Printf("Me: %+v", me)
	}

	answer := 0
	if me.x < 0 {
		answer = answer + me.x*-1
	} else {
		answer = answer + me.x
	}
	if me.y < 0 {
		answer = answer + me.y*-1
	} else {
		answer = answer + me.y
	}

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	type player struct {
		direction int // 0,1,2,3 = N,E,S,W
		x         int
		y         int
		visited   map[string]int
	}
	// Read input file
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	// Read input into list of instructions
	var instructions []string
	instructions = strings.Split(string(f), ", ")
	logger.Print(instructions)

	// Initialize starting pos
	me := player{direction: 0, x: 0, y: 0, visited: map[string]int{"0,0": 1}}
	logger.Printf("Me: %+v", me)

	// Process each instruction
	for i, inst := range instructions {
		logger.Printf("%d: Moving %s", i, inst)

		// Rotate ME
		rotation := string(inst[0])
		switch rotation {
		case "R":
			me.direction = (me.direction + 1) % 4
		case "L":
			me.direction = ((me.direction-1)%4 + 4) % 4
		default:
			logger.Fatalf("Unknown rotation found in instruction: %v", inst)
		}

		// Move ME 1 block at a time so we can track every coord visited
		numStr := inst[1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			logger.Fatal(err)
		}
		for range num {
			switch me.direction {
			case 0:
				me.y++
			case 1:
				me.x++
			case 2:
				me.y--
			case 3:
				me.x--
			default:
				logger.Fatalf("Unknown player direction: %d", me.direction)
			}
			coord := fmt.Sprintf("%d,%d", me.x, me.y)
			me.visited[coord]++
			if me.visited[coord] > 1 {
				answer := 0
				if me.x < 0 {
					answer = answer + me.x*-1
				} else {
					answer = answer + me.x
				}
				if me.y < 0 {
					answer = answer + me.y*-1
				} else {
					answer = answer + me.y
				}
				return answer
			}
		}
		logger.Printf("Me: %+v", me)
	}

	return 0 // This means we didn't find a coord that we visited twice
}
