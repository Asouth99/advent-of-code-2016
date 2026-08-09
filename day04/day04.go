package day04

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day04/input.txt"
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

// A room is real (not a decoy) if the checksum is the five most common letters in the encrypted name, in order, with ties broken by alphabetization.
func isRoomReal(name string, checksum string, logger *log.Logger) bool {
	charsMap := map[string]int{}
	// Loop through chars in name. Keep a map of number of times char is seen
	for _, char := range name {
		if char == '-' {
			continue
		}
		charsMap[string(char)]++
	}
	logger.Printf("Chars: %+v", charsMap)

	// Sort map by num seen. If tie, then sort by alphabetization.
	letters := []string{}
	for k := range charsMap {
		letters = append(letters, k)
	}
	slices.SortFunc(letters, func(a, b string) int {
		if charsMap[a] > charsMap[b] {
			return -1
		} else if charsMap[a] < charsMap[b] {
			return 1
		}
		return strings.Compare(a, b)
	})
	logger.Printf("Sorted Chars: %v", letters)

	checksumAnswer := ""
	for i := 0; i < 5; i++ {
		checksumAnswer += letters[i]
	}

	return checksum == checksumAnswer
}

//ex: aaaaa-bbb-z-y-x-123[abxyz]
// encrypted name: aaaaa-bbb-z-y-x (lowercase letters seperated by dashes)
// sector ID: 123 (numbers)
// checksum: abxyz (letters in square brackets)

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	answer := 0

	line_re := regexp.MustCompile(`^(?P<name>[a-z-]+)-(?P<id>[0-9]+)\[(?P<checksum>.+)\]$`)

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()

		matches := line_re.FindStringSubmatch(line)
		if len(matches) != 4 {
			logger.Fatalf("Could not parse line: %s\nMatches: %v", line, matches)
		}
		nameIdx := line_re.SubexpIndex("name")
		name := matches[nameIdx]
		idIdx := line_re.SubexpIndex("id")
		idStr := matches[idIdx]
		checksumIdx := line_re.SubexpIndex("checksum")
		checksum := matches[checksumIdx]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("Processing Room: {Encrypted Name: %s, Sector ID: %s, Checksum: %s}", name, idStr, checksum)

		// Check if room is real
		if isRoomReal(name, checksum, logger) {
			logger.Print("Room is real")
			answer += id
		} else {
			logger.Print("Room is NOT real")
		}
	}
	return answer
}

func decryptRoom(name string, id int, logger *log.Logger) string {
	decryptedName := ""

	for _, char := range name {
		if char == '-' {
			if id%2 == 0 {
				decryptedName += "-"
			} else {
				decryptedName += " "
			}
		} else {
			decryptedName += string(rune((int(char)-97+id)%26 + 97))
		}
	}
	return decryptedName
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	line_re := regexp.MustCompile(`^(?P<name>[a-z-]+)-(?P<id>[0-9]+)\[(?P<checksum>.+)\]$`)
	decryptedRooms := []string{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()

		matches := line_re.FindStringSubmatch(line)
		if len(matches) != 4 {
			logger.Fatalf("Could not parse line: %s\nMatches: %v", line, matches)
		}
		nameIdx := line_re.SubexpIndex("name")
		name := matches[nameIdx]
		idIdx := line_re.SubexpIndex("id")
		idStr := matches[idIdx]
		// checksumIdx := line_re.SubexpIndex("checksum")
		// checksum := matches[checksumIdx]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("Decrypting Room: {Encrypted Name: %s, Sector ID: %d}", name, id)
		decryptedName := decryptRoom(name, id, logger)
		decryptedRooms = append(decryptedRooms, decryptedName)
		if decryptedName == "northpole-object-storage" {
			return idStr
		}
	}

	for _, room := range decryptedRooms {
		logger.Print(room)
	}

	answer := decryptedRooms[0]
	return answer
}
