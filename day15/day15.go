package day15

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day15/input.txt"
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

type disc struct {
	id            int
	positions     int
	startPosition int
}

func getCapsule(time int, discs []disc) bool {
	for i := range discs {
		disc := discs[i]
		timeAtDisc := time + i + 1
		discPos := (disc.startPosition + timeAtDisc) % disc.positions
		if discPos != 0 {
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

	// Read input discs
	discs := []disc{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)

		id, _ := strconv.Atoi(strings.TrimPrefix(lineSplit[1], "#"))
		positions, _ := strconv.Atoi(lineSplit[3])
		startPos, _ := strconv.Atoi(strings.TrimSuffix(lineSplit[11], "."))

		discs = append(discs, disc{id: id, positions: positions, startPosition: startPos})
	}

	// Print discs
	for i := range discs {
		logger.Printf("Disc #%d, Position: %d/%d", discs[i].id, discs[i].startPosition, discs[i].positions)
	}

	answer := 0
	time := 0
	for true {
		if getCapsule(time, discs) {
			answer = time
			break
		}
		time++
	}

	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read input discs
	discs := []disc{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)

		id, _ := strconv.Atoi(strings.TrimPrefix(lineSplit[1], "#"))
		positions, _ := strconv.Atoi(lineSplit[3])
		startPos, _ := strconv.Atoi(strings.TrimSuffix(lineSplit[11], "."))

		discs = append(discs, disc{id: id, positions: positions, startPosition: startPos})
	}
	// Add new disc
	discs = append(discs, disc{id: len(discs) + 1, positions: 11, startPosition: 0})

	// Print discs
	for i := range discs {
		logger.Printf("Disc #%d, Position: %d/%d", discs[i].id, discs[i].startPosition, discs[i].positions)
	}

	answer := 0
	time := 0
	for true {
		if getCapsule(time, discs) {
			answer = time
			break
		}
		time++
	}

	return answer
}
