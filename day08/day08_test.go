package day08_test

import (
	"aoc2016/day08"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

// Define the structure for a single test case
type testCase struct {
	name          string
	inputFilePath string
	expectedP1    any // Expected answer for Part 1
	expectedP2    any // Expected answer for Part 2
}

var tests = []testCase{
	{
		name:          "Example 1",
		inputFilePath: "example_1.txt",
		expectedP1:    6,
		expectedP2:    0,
	},
}

var testLogger *log.Logger

func TestMain(m *testing.M) {
	verbosePtr := flag.Bool("log", false, "Enable verbose logging")
	flag.Parse()
	verbose := *verbosePtr

	if verbose {
		testLogger = log.New(os.Stderr, "[TEST LOG] ", log.Ltime|log.Lshortfile)
		fmt.Println("--- Custom Logging Enabled ---")
	} else {
		testLogger = log.New(io.Discard, "", 0)
	}

	os.Exit(m.Run())
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tests {
		if tc.expectedP1 == 0 {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			p1, err := day08.Solve(1, testLogger, tc.inputFilePath)
			if err != nil {
				t.Fatalf("❌ Part 1 failed to run: %v\n", err)
				return
			}
			if p1 != tc.expectedP1 {
				t.Errorf("Part 1 failed. Expected: %d, Got: %d", tc.expectedP1, p1)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tests {
		if tc.expectedP2 == 0 {
			continue
		}
		t.Run(tc.name, func(t *testing.T) {
			p2, err := day08.Solve(2, testLogger, tc.inputFilePath)
			if err != nil {
				t.Fatalf("❌ Part 2 failed to run: %v\n", err)
				return
			}
			if p2 != tc.expectedP2 {
				t.Errorf("Part 2 failed. Expected: %d, Got: %d", tc.expectedP2, p2)
			}
		})
	}
}
