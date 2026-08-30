package day21

import (
	"bufio"
	"errors"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day21/input.txt"
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

// swap position X with position Y means that the letters at indexes X and Y (counting from 0) should be swapped.
// swap letter X with letter Y means that the letters X and Y should be swapped (regardless of where they appear in the string).
// rotate left/right X steps means that the whole string should be rotated; for example, one right rotation would turn abcd into dabc.
// rotate based on position of letter X means that the whole string should be rotated to the right based on the index of letter X (counting from 0) as determined before this instruction does any rotations. Once the index is determined, rotate the string to the right one time, plus a number of times equal to that index, plus one additional time if the index was at least 4.
// reverse positions X through Y means that the span of letters at indexes X through Y (including the letters at X and Y) should be reversed in order.
// move position X to position Y means that the letter which is at index X should be removed from the string, then inserted such that it ends up at index Y.
func swapLetters(str []byte, charX byte, charY byte) {
	for i := range str {
		if str[i] == charX {
			str[i] = charY
		} else if str[i] == charY {
			str[i] = charX
		}
	}
}
func swap(str []byte, posX int, posY int) {
	str[posX], str[posY] = str[posY], str[posX]
}
func rotate(str []byte, amount int, left bool) {
	copy := slices.Clone(str)
	if left {
		for i := range str {
			nextIdx := ((i+amount)%len(str) + len(str)) % len(str)
			str[i] = copy[nextIdx]
		}
	} else {
		for i := range str {
			nextIdx := ((i-amount)%len(str) + len(str)) % len(str)
			str[i] = copy[nextIdx]
		}
	}
}

func reverse(str []byte, posX int, posY int) {
	// edcba
	// reverse positions 0 through 4
	// abcde
	l := posX
	r := posY
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
}
func move(str []byte, posX int, posY int) {
	if posX == posY {
		return
	}
	char := str[posX]
	if posX < posY {
		copy(str[posX:posY], str[posX+1:posY+1])
	} else {
		copy(str[posY+1:posX+1], str[posY:posX])
	}
	str[posY] = char
}

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	var input []byte
	instructions := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i == 0 {
			input = []byte(line)
		} else {
			instructions = append(instructions, strings.Fields(line))
		}
	}

	// Print inputs
	logger.Printf("Scrambling <%s>", input)
	for i := range instructions {
		logger.Print(instructions[i])
	}
	logger.Print()

	// Go through every instruction
	for _, inst := range instructions {
		logger.Print(inst)
		str := string(input) + " -> "

		switch inst[0] {
		case "swap":
			if inst[1] == "position" {
				posX, _ := strconv.Atoi(inst[2])
				posY, _ := strconv.Atoi(inst[5])
				swap(input, posX, posY)
			} else {
				swapLetters(input, byte(inst[2][0]), byte(inst[5][0]))
			}
		case "reverse":
			posX, _ := strconv.Atoi(inst[2])
			posY, _ := strconv.Atoi(inst[4])
			reverse(input, posX, posY)
		case "rotate":
			if inst[1] == "right" {
				amount, _ := strconv.Atoi(inst[2])
				rotate(input, amount, false)
			} else if inst[1] == "left" {
				amount, _ := strconv.Atoi(inst[2])
				rotate(input, amount, true)
			} else {
				// Get index of char
				char := byte(inst[6][0])
				idx := slices.Index(input, char)
				amount := idx + 1
				if idx >= 4 {
					amount++
				}
				rotate(input, amount, false)
			}
		case "move":
			posX, _ := strconv.Atoi(inst[2])
			posY, _ := strconv.Atoi(inst[5])
			move(input, posX, posY)
		default:
			logger.Fatalf("Unknown instruction %s", inst)
		}

		str += string(input)
		logger.Print(str)
	}

	answer := string(input)
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	var input []byte
	instructions := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i == 0 {
			input = []byte(line)
		} else {
			instructions = append(instructions, strings.Fields(line))
		}
	}
	// Overwite input
	if strings.Contains(inputFile, "input") {
		input = []byte{'f', 'b', 'g', 'd', 'c', 'e', 'a', 'h'}
		logger.Print("Overriding input")
	}

	// Lookup table for rotationsByIndex
	rotationsByIndex := []int{}
	for i := range len(input) {
		newIdx := i + i + 1
		if i >= 4 {
			newIdx++
		}
		newIdx = newIdx % len(input)
		rotationsByIndex = append(rotationsByIndex, newIdx)
	}
	logger.Print(rotationsByIndex)

	// Lookup table for unrotationsByIndex
	unrotationsByIndex := []int{}
	for i := range rotationsByIndex {
		oldIdx := 0
		// Find newIdx in rotationsByIndex
		for j, idx := range rotationsByIndex {
			if idx == i {
				oldIdx = j
				// break
			}
		}
		rot := ((i-oldIdx)%len(input) + len(input)) % len(input)
		unrotationsByIndex = append(unrotationsByIndex, rot)
	}
	logger.Print(unrotationsByIndex)

	// Print inputs
	// logger.Printf("Unscrambling <%s>", input)
	// for i := range instructions {
	// 	logger.Print(instructions[i])
	// }
	// logger.Print()

	// Go through every instruction backwards
	for i := len(instructions) - 1; i >= 0; i-- {
		inst := instructions[i]
		logger.Print(inst)
		str := string(input) + " -> "

		switch inst[0] {
		case "swap":
			if inst[1] == "position" {
				posX, _ := strconv.Atoi(inst[2])
				posY, _ := strconv.Atoi(inst[5])
				swap(input, posX, posY)
			} else {
				swapLetters(input, byte(inst[2][0]), byte(inst[5][0]))
			}
		case "reverse":
			posX, _ := strconv.Atoi(inst[2])
			posY, _ := strconv.Atoi(inst[4])
			reverse(input, posX, posY)
		case "rotate":
			if inst[1] == "right" {
				amount, _ := strconv.Atoi(inst[2])
				rotate(input, amount, true)
			} else if inst[1] == "left" {
				amount, _ := strconv.Atoi(inst[2])
				rotate(input, amount, false)
			} else {
				// Get index of char
				// TODO: FIND OUT HOW MUCH TO ROTATE BY
				char := byte(inst[6][0])
				idx := slices.Index(input, char)
				rotate(input, unrotationsByIndex[idx], true)
			}
		case "move":
			posX, _ := strconv.Atoi(inst[2])
			posY, _ := strconv.Atoi(inst[5])
			move(input, posY, posX)
		default:
			logger.Fatalf("Unknown instruction %s", inst)
		}

		str += string(input)
		logger.Print(str)
	}

	answer := string(input)
	return answer
}
