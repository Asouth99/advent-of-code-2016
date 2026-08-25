package day16

import (
	"errors"
	"log"
	"os"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day16/input.txt"
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

// Call the data you have at this point "a".
// Make a copy of "a"; call this copy "b".
// Reverse the order of the characters in "b".
// In "b", replace all instances of 0 with 1 and all 1s with 0.
// The resulting data is "a", then a single 0, then "b".
func generateString(initialState string, length int) string {
	str := initialState
	for len(str) < length {
		a := str
		str = a + "0"
		for i := len(a) - 1; i >= 0; i-- {
			char := a[i]
			switch char {
			case '1':
				str += "0"
			case '0':
				str += "1"
			}
		}
	}
	return str[:length]
}

// Consider each pair: 11, 00, 10, 11, 01, 00.
// These are same, same, different, same, different, same, producing 110101.
// The resulting string has length 6, which is even, so we repeat the process.
// The pairs are 11 (same), 01 (different), 01 (different).
// This produces the checksum 100, which has an odd length, so we stop.
func checksum(str string) string {
	checksum := str
	for len(checksum)%2 == 0 {
		prev := checksum
		new := ""
		for i := 0; i < len(prev)-1; i += 2 {
			if prev[i] == prev[i+1] {
				new += "1"
			} else {
				new += "0"
			}
		}
		checksum = new
	}
	return checksum
}

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	var diskLength int
	if strings.HasPrefix(inputFile, "example") {
		diskLength = 20
	} else {
		diskLength = 272
	}

	logger.Printf("Generating string of length %d with initial state %s", diskLength, string(f))
	data := generateString(string(f), diskLength)
	logger.Printf("Calculating checksum of %s", data)
	checksum := checksum(data)
	logger.Print(checksum)

	answer := checksum
	return answer
}

// Optimised to use a string builder
func generateString2(initialState string, length int) string {
	strBuilder := strings.Builder{}
	strBuilder.Grow(length)
	strBuilder.WriteString(initialState)
	for len(strBuilder.String()) < length {
		a := strBuilder.String()
		strBuilder.WriteRune('0')
		for i := len(a) - 1; i >= 0; i-- {
			char := a[i]
			switch char {
			case '1':
				strBuilder.WriteRune('0')
			case '0':
				strBuilder.WriteRune('1')
			}
		}
	}
	return strBuilder.String()[:length]
}

// Optimised to use a string builder
func checksum2(str string) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(str)
	for len(strBuilder.String())%2 == 0 {
		prev := strBuilder.String()
		strBuilder.Reset()
		strBuilder.Grow(len(prev) / 2)
		for i := 0; i < len(prev)-1; i += 2 {
			if prev[i] == prev[i+1] {
				strBuilder.WriteRune('1')
			} else {
				strBuilder.WriteRune('0')
			}
		}
	}
	return strBuilder.String()
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}

	var diskLength int
	if strings.HasPrefix(inputFile, "example") {
		diskLength = 20
	} else {
		diskLength = 35651584
	}

	logger.Printf("Generating string of length %d with initial state %s", diskLength, string(f))
	data := generateString2(string(f), diskLength)
	logger.Printf("Calculating checksum of %s", data)
	checksum := checksum2(data)
	logger.Printf("Checksum: %s", checksum)

	answer := checksum
	return answer
}
