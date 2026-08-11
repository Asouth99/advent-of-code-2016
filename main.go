package main

import (
	// Placeholder for ./template.sh to modify
	// ADD IMPORT HERE
    "aoc2016/day08"
    "aoc2016/day07"
    "aoc2016/day06"
    "aoc2016/day05"
    "aoc2016/day04"
    "aoc2016/day03"
    "aoc2016/day02"
	"aoc2016/day01"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// A type alias for our standardized Solve function signature
type Solver func(int, *log.Logger, ...string) (any, error)

// A map to associate the day number (int) with the corresponding Solve function
var solutions = map[int]Solver{
	// Placeholder for ./template.sh to modify
	// ADD SOLUTION HERE
    8: day08.Solve,
    7: day07.Solve,
    6: day06.Solve,
    5: day05.Solve,
    4: day04.Solve,
    3: day03.Solve,
    2: day02.Solve,
	1: day01.Solve,
}

// Global logger that will be used across the application for verbose messages.
var verboseLogger *log.Logger

func main() {
	// Parse command line flags
	dayPtr := flag.Int("day", 0, "The Advent of Code day number to run (1-25)")
	partPtr := flag.Int("part", 0, "The part to run 1 or 2")
	verbosePtr := flag.Bool("v", false, "Enable verbose logging")
	flag.Parse()
	day := *dayPtr
	part := *partPtr
	verbose := *verbosePtr

	if verbose {
		verboseLogger = log.New(os.Stderr, "[VERBOSE] ", log.Ltime|log.Lshortfile)
	} else {
		verboseLogger = log.New(io.Discard, "", 0)
	}

	if day == 0 {
		fmt.Println("Usage: go run main.go --day=<number>")
		fmt.Println("Example: go run main.go --day=1")
		return
	}

	fmt.Println("--- Advent of Code 2025 ---")

	if part != 0 {
		fmt.Printf("Running solution for Day %02d Part %d\n", day, part)
		runSolution(day, part)
	} else {
		fmt.Printf("Running solutions for Day %02d\n", day)
		runAllSolutions(day)
	}

}

func runAllSolutions(day int) {
	solveFunc, exists := solutions[day]
	if !exists {
		fmt.Printf("Error: Solution for Day %02d not found.\n", day)
		return
	}
	p1, err := solveFunc(1, verboseLogger)
	if err != nil {
		fmt.Printf("❌ Day %02d Part 1 failed to run: %v\n", day, err)
		return
	}
	p2, err := solveFunc(2, verboseLogger)
	if err != nil {
		fmt.Printf("❌ Day %02d Part 2 failed to run: %v\n", day, err)
		return
	}
	fmt.Printf("Day %02d: Part 1 = %d | Part 2 = %d\n", day, p1, p2)
}

func runSolution(day int, part int) {
	solveFunc, exists := solutions[day]
	if !exists {
		fmt.Printf("Error: Solution for Day %02d not found.\n", day)
		return
	}
	answer, err := solveFunc(part, verboseLogger)
	if err != nil {
		fmt.Printf("❌ Day %02d Part %d failed to run: %v\n", day, part, err)
		return
	}
	fmt.Printf("Day %02d: Part %d = %v\n", day, part, answer)
}
