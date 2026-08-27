package day17

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"math"
	"os"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day17/input.txt"
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

func getDoorStates(str string) [4]bool {
	hashBytes := md5.Sum([]byte(str))
	hash := hex.EncodeToString(hashBytes[:])
	return [4]bool{
		strings.ContainsRune("bcdef", rune(hash[0])), // Up
		strings.ContainsRune("bcdef", rune(hash[1])), // Down
		strings.ContainsRune("bcdef", rune(hash[2])), // Left
		strings.ContainsRune("bcdef", rune(hash[3])), // Right
	}
}

func findShortestPath(max [2]int, pos [2]int, target [2]int, minPathLength *int, path string, passcode string) string {
	pathLength := len(path)

	// Exit early if path is longer than known min
	if pathLength >= *minPathLength {
		return ""
	}

	// Check if we are at the end
	if pos[0] == target[0] && pos[1] == target[1] {
		if pathLength < *minPathLength {
			*minPathLength = pathLength
			return path
		}
	}

	// Check every direction from current pos
	dirs := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	dirChars := [4]string{"U", "D", "L", "R"}
	doorStates := getDoorStates(passcode + path)
	bestPath := ""
	for i, dir := range dirs {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if doorStates[i] && nextPos[0] >= 0 && nextPos[1] >= 0 && nextPos[0] < max[0] && nextPos[1] < max[1] {
			result := findShortestPath(max, nextPos, target, minPathLength, path+dirChars[i], passcode)
			if result != "" {
				bestPath = result
			}
		}
	}

	return bestPath
}

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	passcode := string(f)

	max := [2]int{4, 4}      // max x,y
	startPos := [2]int{0, 0} // pos x,y
	target := [2]int{max[0] - 1, max[1] - 1}
	pathLength := math.MaxInt64

	path := findShortestPath(max, startPos, target, &pathLength, "", passcode)

	answer := path
	return answer
}

func findLongestPath(max [2]int, pos [2]int, target [2]int, maxPathLength *int, path string, passcode string) {
	pathLength := len(path)

	// Check if we are at the end
	if pos[0] == target[0] && pos[1] == target[1] {
		if pathLength > *maxPathLength {
			*maxPathLength = pathLength
		}
		return
	}

	// Check every direction from current pos
	dirs := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	dirChars := [4]string{"U", "D", "L", "R"}
	doorStates := getDoorStates(passcode + path)
	for i, dir := range dirs {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if doorStates[i] && nextPos[0] >= 0 && nextPos[1] >= 0 && nextPos[0] < max[0] && nextPos[1] < max[1] {
			findLongestPath(max, nextPos, target, maxPathLength, path+dirChars[i], passcode)
		}
	}

	return
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	passcode := string(f)

	max := [2]int{4, 4}      // max x,y
	startPos := [2]int{0, 0} // pos x,y
	target := [2]int{max[0] - 1, max[1] - 1}
	pathLength := 0

	findLongestPath(max, startPos, target, &pathLength, "", passcode)

	answer := pathLength
	return answer
}
