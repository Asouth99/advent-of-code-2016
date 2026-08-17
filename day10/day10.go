package day10

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day10/input.txt"
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

type bot struct {
	low  int
	high int
}

func (b *bot) addValue(v int) {
	// Bot is full
	if b.low != 0 && b.high != 0 {
		return
	}

	// Empty bot
	if b.low == 0 {
		b.low = v
		return
	}

	// Sort vals
	if v < b.low {
		b.high = b.low
		b.low = v
	} else {
		b.high = v
	}
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	var answer_bot bot
	if strings.HasPrefix(inputFile, "example") {
		answer_bot = bot{high: 5, low: 2}
	} else {
		answer_bot = bot{high: 61, low: 17}
	}
	logger.Printf("Looking for bot %+v", answer_bot)

	answer := 0
	bots := map[int]*bot{}
	outputs := map[int]int{}
	instructions := map[int][]string{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		logger.Printf("%d : %s", i, line)

		lineSplit := strings.Fields(line)
		switch lineSplit[0] {
		case "value":
			val, _ := strconv.Atoi(lineSplit[1])
			botNum, _ := strconv.Atoi(lineSplit[5])
			if b, ok := bots[botNum]; ok {
				b.addValue(val)
			} else {
				b := bot{}
				b.addValue(val)
				bots[botNum] = &b
			}
		case "bot":
			botNum, _ := strconv.Atoi(lineSplit[1])
			instructions[botNum] = lineSplit
			lowDest := lineSplit[5] // bot or output
			lowDestNum, _ := strconv.Atoi(lineSplit[6])
			highDest := lineSplit[10] // bot or output
			highDestNum, _ := strconv.Atoi(lineSplit[11])

			if lowDest == "output" {
				if _, ok := outputs[lowDestNum]; !ok {
					outputs[lowDestNum] = 0
				}
			}
			if lowDest == "bot" {
				if _, ok := bots[lowDestNum]; !ok {
					bots[lowDestNum] = &bot{}
				}
			}
			if highDest == "output" {
				if _, ok := outputs[highDestNum]; !ok {
					outputs[highDestNum] = 0
				}
			}
			if highDest == "bot" {
				if _, ok := bots[highDestNum]; !ok {
					bots[highDestNum] = &bot{}
				}
			}
		default:
			logger.Fatalf("Error: Unknown start of line '%s'", line)
		}
	}

	// Print all bots
	logger.Print("--- Bots initial state ---")
	for k, b := range bots {
		logger.Printf("bot %d : %+v", k, *b)
	}

	// Go through all instructions
	for {
		progress := false
		for botNum, bot := range bots {
			// Skip if bot doesn't have both values
			if bot.low == 0 || bot.high == 0 {
				continue
			}

			// Check if it is comparing the values required
			if *bot == answer_bot {
				answer = botNum
				progress = false
				break
			}

			// Process bot
			inst := instructions[botNum]
			lowDest := inst[5] // bot or output
			lowDestNum, _ := strconv.Atoi(inst[6])
			highDest := inst[10] // bot or output
			highDestNum, _ := strconv.Atoi(inst[11])
			if lowDest == "output" {
				outputs[lowDestNum] = bot.low
			} else {
				bots[lowDestNum].addValue(bot.low)
			}
			if highDest == "output" {
				outputs[highDestNum] = bot.high
			} else {
				bots[highDestNum].addValue(bot.high)
			}
			bot.low = 0
			bot.high = 0
			progress = true

		}
		if progress == false {
			break
		}
	}

	if answer == 0 {
		logger.Fatalf("No bot found that matched %+v", answer_bot)
	}
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	var answer_bot bot
	if strings.HasPrefix(inputFile, "example") {
		answer_bot = bot{high: 5, low: 2}
	} else {
		answer_bot = bot{high: 61, low: 17}
	}
	logger.Printf("Looking for bot %+v", answer_bot)

	answer := 0
	bots := map[int]*bot{}
	outputs := map[int]int{}
	instructions := map[int][]string{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		logger.Printf("%d : %s", i, line)

		lineSplit := strings.Fields(line)
		switch lineSplit[0] {
		case "value":
			val, _ := strconv.Atoi(lineSplit[1])
			botNum, _ := strconv.Atoi(lineSplit[5])
			if b, ok := bots[botNum]; ok {
				b.addValue(val)
			} else {
				b := bot{}
				b.addValue(val)
				bots[botNum] = &b
			}
		case "bot":
			botNum, _ := strconv.Atoi(lineSplit[1])
			instructions[botNum] = lineSplit
			lowDest := lineSplit[5] // bot or output
			lowDestNum, _ := strconv.Atoi(lineSplit[6])
			highDest := lineSplit[10] // bot or output
			highDestNum, _ := strconv.Atoi(lineSplit[11])

			if lowDest == "output" {
				if _, ok := outputs[lowDestNum]; !ok {
					outputs[lowDestNum] = 0
				}
			}
			if lowDest == "bot" {
				if _, ok := bots[lowDestNum]; !ok {
					bots[lowDestNum] = &bot{}
				}
			}
			if highDest == "output" {
				if _, ok := outputs[highDestNum]; !ok {
					outputs[highDestNum] = 0
				}
			}
			if highDest == "bot" {
				if _, ok := bots[highDestNum]; !ok {
					bots[highDestNum] = &bot{}
				}
			}
		default:
			logger.Fatalf("Error: Unknown start of line '%s'", line)
		}
	}

	// Print all bots
	logger.Print("--- Bots initial state ---")
	for k, b := range bots {
		logger.Printf("bot %d : %+v", k, *b)
	}

	// Go through all instructions
	for {
		progress := false
		for botNum, bot := range bots {
			// Skip if bot doesn't have both values
			if bot.low == 0 || bot.high == 0 {
				continue
			}

			// Check if it is comparing the values required
			// if *bot == answer_bot {
			// 	answer = botNum
			// 	progress = false
			// 	break
			// }

			// Process bot
			inst := instructions[botNum]
			lowDest := inst[5] // bot or output
			lowDestNum, _ := strconv.Atoi(inst[6])
			highDest := inst[10] // bot or output
			highDestNum, _ := strconv.Atoi(inst[11])
			if lowDest == "output" {
				outputs[lowDestNum] = bot.low
			} else {
				bots[lowDestNum].addValue(bot.low)
			}
			if highDest == "output" {
				outputs[highDestNum] = bot.high
			} else {
				bots[highDestNum].addValue(bot.high)
			}
			bot.low = 0
			bot.high = 0
			progress = true

		}
		if progress == false {
			break
		}
	}

	answer = outputs[0] * outputs[1] * outputs[2]
	return answer
}
