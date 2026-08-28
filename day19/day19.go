package day19

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day19/input.txt"
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

type elf struct {
	id     int
	target *elf
}

func printCircle(head *elf) {
	str := fmt.Sprintf("%d -> ", head.id)
	current := head.target
	for current != head {
		str += fmt.Sprintf("%d -> ", current.id)
		current = current.target
	}

	fmt.Print(str, head.id, "\n")
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	numElfs, err := strconv.Atoi(string(f))
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Finding which elf gets all the presents in a circle of %d", numElfs)

	// Initialise circular linked list
	head := &elf{id: 1}
	current := head
	for i := 2; i <= numElfs; i++ {
		next := &elf{id: i}
		current.target = next
		current = next
	}
	current.target = head

	if strings.HasPrefix(inputFile, "example") {
		printCircle(head)
	}

	// Keep looping until an elf points to itself
	current = head
	for current != current.target {
		current.target = current.target.target
		current = current.target
	}

	answer := current.id
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	numElfs, err := strconv.Atoi(string(f))
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Finding which elf gets all the presents in a circle of %d", numElfs)

	answer := 0
	return answer
}
