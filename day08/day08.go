package day08

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day08/input.txt"
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

func rectPixels(pixels [][]bool, width int, height int) {
	for x := range width {
		for y := range height {
			pixels[x][y] = true
		}
	}
}

func rotatePixelRow(pixels [][]bool, r int, amount int) {
	width := len(pixels)
	// copy the row
	tmp := make([]bool, width)
	for x := range width {
		tmp[x] = pixels[x][r]
	}
	// Update the slice
	for x := range width {
		pixels[x][r] = tmp[((x-amount)%width+width)%width]
	}
}

func rotatePixelColumn(pixels [][]bool, c int, amount int) {
	height := len(pixels[c])
	// copy the row
	tmp := make([]bool, height)
	for y := range height {
		tmp[y] = pixels[c][y]
	}
	// Update the slice
	for y := range height {
		pixels[c][y] = tmp[((y-amount)%height+height)%height]
	}
}

func printPixels(pixels [][]bool, logger *log.Logger) {
	str := "Printing pixels"
	for y := range len(pixels[0]) {
		str += "\n"
		for x := range len(pixels) {
			if pixels[x][y] {
				str += "#"
			} else {
				str += "."
			}
		}
	}
	logger.Print(str)
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	MAX_WIDTH := 50
	MAX_HEIGHT := 6
	if strings.HasPrefix(inputFile, "example") {
		MAX_WIDTH = 7
		MAX_HEIGHT = 3
	}

	pixels := make([][]bool, MAX_WIDTH)
	for i := range pixels {
		pixels[i] = make([]bool, MAX_HEIGHT)
	}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		printPixels(pixels, logger)

		line := scanner.Text()
		logger.Printf("%s", line)
		lineSplit := strings.Fields(line)
		switch lineSplit[0] {
		case "rect":
			wh := strings.Split(lineSplit[1], "x")
			w, err := strconv.Atoi(wh[0])
			if err != nil {
				logger.Fatal(err)
			}
			h, err := strconv.Atoi(wh[1])
			if err != nil {
				logger.Fatal(err)
			}
			rectPixels(pixels, w, h)
		case "rotate":
			r_c, err := strconv.Atoi(strings.Split(lineSplit[2], "=")[1])
			if err != nil {
				logger.Fatal(err)
			}
			amount, err := strconv.Atoi(lineSplit[4])
			if err != nil {
				logger.Fatal(err)
			}
			if lineSplit[1] == "column" {
				rotatePixelColumn(pixels, r_c, amount)
			} else {
				rotatePixelRow(pixels, r_c, amount)
			}
		default:
			logger.Fatalf("Error parsing input line: %s", line)
		}
	}

	logger.Print("--- Final Pixel output ---")
	printPixels(pixels, logger)

	answer := 0

	for x := range len(pixels) {
		for y := range len(pixels[x]) {
			if pixels[x][y] {
				answer++
			}
		}
	}
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	MAX_WIDTH := 50
	MAX_HEIGHT := 6
	if strings.HasPrefix(inputFile, "example") {
		MAX_WIDTH = 7
		MAX_HEIGHT = 3
	}

	pixels := make([][]bool, MAX_WIDTH)
	for i := range pixels {
		pixels[i] = make([]bool, MAX_HEIGHT)
	}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		printPixels(pixels, logger)

		line := scanner.Text()
		logger.Printf("%s", line)
		lineSplit := strings.Fields(line)
		switch lineSplit[0] {
		case "rect":
			wh := strings.Split(lineSplit[1], "x")
			w, err := strconv.Atoi(wh[0])
			if err != nil {
				logger.Fatal(err)
			}
			h, err := strconv.Atoi(wh[1])
			if err != nil {
				logger.Fatal(err)
			}
			rectPixels(pixels, w, h)
		case "rotate":
			r_c, err := strconv.Atoi(strings.Split(lineSplit[2], "=")[1])
			if err != nil {
				logger.Fatal(err)
			}
			amount, err := strconv.Atoi(lineSplit[4])
			if err != nil {
				logger.Fatal(err)
			}
			if lineSplit[1] == "column" {
				rotatePixelColumn(pixels, r_c, amount)
			} else {
				rotatePixelRow(pixels, r_c, amount)
			}
		default:
			logger.Fatalf("Error parsing input line: %s", line)
		}
	}

	logger.Print("--- Final Pixel output ---")
	printPixels(pixels, logger)

	answer := 0

	for x := range len(pixels) {
		for y := range len(pixels[x]) {
			if pixels[x][y] {
				answer++
			}
		}
	}
	return answer
}
