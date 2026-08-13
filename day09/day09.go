package day09

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day09/input.txt"
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
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	// decompressed := []byte{}
	var decompressedBuilder strings.Builder
	for i := 0; i < len(f); i++ {
		char := f[i]
		if char == '(' { // Start of marker
			// Find idx of end of marker
			idxEnd := bytes.IndexByte(f[i:], ')')
			if idxEnd == -1 {
				logger.Fatal("No closing ) for marker")
			}
			idxEnd += i
			// Get marker values
			marker := strings.Split(string(f[i+1:idxEnd]), "x")
			chars, err := strconv.Atoi(marker[0])
			if err != nil {
				logger.Fatal(err)
			}
			repeat, err := strconv.Atoi(marker[1])
			if err != nil {
				logger.Fatal(err)
			}
			logger.Printf("Found marker %d, %d", chars, repeat)

			charsToRepeat := f[idxEnd+1 : idxEnd+1+chars]
			logger.Printf("Repeating '%s' %d times", string(charsToRepeat), repeat)

			for range repeat {
				decompressedBuilder.Write(charsToRepeat)
			}

			// Update outer loop to skip over characters already repeated by marker
			i = idxEnd + chars

			// Add chars to decompressed string
		} else {
			decompressedBuilder.WriteByte(char)
		}
	}

	decompressedString := decompressedBuilder.String()
	logger.Printf("len=%d | %s", len(decompressedString), decompressedString)

	answer := len(decompressedString)
	return answer
}

// ex:
// (27x12)(20x12)(13x14)(7x10)(1x12)A
// (20x12)(13x14)(7x10)(1x12)A * 12
// (13x14)(7x10)(1x12)A * 12 * 12
// (7x10)(1x12)A * 14 * 12 * 12
// (1x12)A * 10 * 14 * 12 * 12
// A * 12 * 10 * 14 * 12 * 12
// 1 * 12 * 10 * 14 * 12 * 12 = 241920 chars

// ex:
// aa(3x3)XYZX(8x2)(3x3)ABCY(27x12)(20x12)(13x14)(7x10)(1x12)A(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN
// (3x3)XYZ X (8x2)(3x3)ABC Y (27x12)(20x12)(13x14)(7x10)(1x12)A (25x3)(3x3)ABC(2x3)XY(5x2)PQRST X (18x9)(3x2)TWO(5x7)SEVEN
// ex:
// (25x3)(3x3)ABC(2x3)XY(5x2)PQRST
// 3 * 3 * ABC + 3 * 3 * XY + 3 * 2 * PQRST
// 3 * 3 * 3 + 3 * 3 * 2 + 3 * 2 * 5 = 75

// DecompressedLength recursively calculates the length of decompressed data (Part 2)
func DecompressedLength(input []byte, logger *log.Logger) int {
	length := 0
	i := 0

	for i < len(input) {
		if input[i] == '(' {
			// Find closing parenthesis
			closeIdx := bytes.IndexByte(input[i:], ')')
			if closeIdx == -1 {
				log.Fatal("Unclosed marker")
			}
			closeIdx += i // Adjust offset relative to full slice

			// Extract marker numbers "AxB"
			markerStr := string(input[i+1 : closeIdx])
			parts := strings.Split(markerStr, "x")
			numChars, err := strconv.Atoi(parts[0])
			if err != nil {
				logger.Fatal(err)
			}
			repeat, err := strconv.Atoi(parts[1])
			if err != nil {
				logger.Fatal(err)
			}

			// Get the chars that this marker repeats
			dataStart := closeIdx + 1
			dataEnd := dataStart + numChars
			repeatedData := input[dataStart:dataEnd]

			// Recursively decompress the chars and multiply its decompressed length
			length += DecompressedLength(repeatedData, logger) * repeat

			// Move i to the end of the marker + chars
			i = dataEnd
		} else {
			// Normal character, counts as 1 byte
			length++
			i++
		}
	}

	return length
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	answer := DecompressedLength(f, logger)
	return answer
}
