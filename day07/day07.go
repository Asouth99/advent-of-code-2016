package day07

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day07/input.txt"
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

// An IP supports TLS if it has an Autonomous Bridge Bypass Annotation, or ABBA.
// An ABBA is any four-character sequence which consists of a pair of two different characters followed by the reverse of that pair,
// ex: xyyx, abba.
// However, the IP also must not have an ABBA within any hypernet sequences, which are contained by square brackets.
func supportsTls(ip string) bool {
	inSquareBrackets := false
	abbaSequenceFound := false
	for i, char := range ip {
		if i > len(ip)-4 {
			break
		}

		if char == '[' {
			inSquareBrackets = true
		} else if char == ']' {
			inSquareBrackets = false
		}

		// Check if ABBA sequence
		if ip[i] != ip[i+1] && ip[i] == ip[i+3] && ip[i+1] == ip[i+2] {
			if inSquareBrackets == true {
				return false
			} else {
				abbaSequenceFound = true
			}
		}

	}
	return abbaSequenceFound
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		logger.Printf("Checking if IP %s supports TLS", line)
		supportsTls := supportsTls(line)
		if supportsTls {
			logger.Printf("Supports TLS!")
		}

		if supportsTls {
			answer++
		}

	}
	return answer
}

// ANSI just to make the terminal output fancy
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// ABA, anywhere in the outside square bracketed sections
// and a corresponding BAB, anywhere in the square bracket sequences.
// An ABA is any three-character sequence which consists of the same character twice with a different character between them
// ex: xyx, aba.
// A corresponding BAB is the same characters but in reversed positions:
// ex: yxy and bab
// ex: aba[bab]xyz has aba outside square brakcets and corresponding bab in square brackets
func supportsSsl(supers []string, hypers []string) bool {
	for _, super := range supers {
		for i := range super {
			if i > len(super)-3 {
				break
			}

			// Check if ABA sequence
			if super[i] != super[i+1] && super[i] == super[i+2] {
				// Check if any hyper contains BAB
				bab := string(super[i+1]) + string(super[i]) + string(super[i+1])
				for _, hyper := range hypers {
					if strings.Contains(hyper, bab) {
						return true
					}
				}
			}
		}
	}
	return false
}

func parseIpv7(ip string) ([]string, []string) {
	supers := []string{}
	hypers := []string{}
	idx := 0
	for i, char := range ip {
		switch char {
		case '[':
			supers = append(supers, ip[idx:i])
			idx = i + 1
		case ']':
			hypers = append(hypers, ip[idx:i])
			idx = i + 1
		}
	}
	supers = append(supers, ip[idx:]) // Add last super sequence

	return supers, hypers
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		logger.Printf("Checking if IP "+Blue+"%s"+Reset+" supports SSL", line)

		supers, hypers := parseIpv7(line)

		logger.Printf("SuperSequences: %v | HyperSequences: %v", supers, hypers)

		supportsSsl := supportsSsl(supers, hypers)
		if supportsSsl {
			logger.Print(Green + "Supports SSL!" + Reset)
		} else {
			logger.Print(Red + "Doesn't support SSL!" + Reset)
		}

		if supportsSsl {
			answer++
		}

	}
	return answer
}
