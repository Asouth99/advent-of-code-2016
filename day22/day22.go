package day22

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day22/input.txt"
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

type node struct {
	name  string
	size  int
	used  int
	avail int
	use   int
	x     int
	y     int
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	nodes := []node{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i >= 2 {
			lineSplit := strings.Fields(line)
			nums := []int{}
			for j := 1; j < len(lineSplit); j++ {
				num, err := strconv.Atoi(lineSplit[j][:len(lineSplit[j])-1])
				if err != nil {
					logger.Fatal(err)
				}
				nums = append(nums, num)
			}
			n := node{
				name:  lineSplit[0],
				size:  nums[0],
				used:  nums[1],
				avail: nums[2],
				use:   nums[3],
			}
			nodes = append(nodes, n)
		}
	}

	// Print input nodes
	logger.Print("-------------------------- Nodes --------------------------")
	for _, n := range nodes {
		logger.Printf("%+v", n)
	}

	// Loop through all nodes, check every other node
	nodePairs := [][2]node{}
	for a := 0; a < len(nodes); a++ {
		nodeA := nodes[a]
		if nodeA.used <= 0 {
			continue
		}
		for b := 0; b < len(nodes); b++ {
			if a == b {
				continue
			}
			nodeB := nodes[b]

			if nodeA.used <= nodeB.avail {
				// logger.Printf("found node pair\n%+v\n%+v", nodeA, nodeB)
				nodePairs = append(nodePairs, [2]node{nodeA, nodeB})
			}
		}
	}

	answer := len(nodePairs)
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	nodes := []node{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if i >= 2 {
			lineSplit := strings.Fields(line)
			nameSplit := strings.Split(lineSplit[0], "-")
			nums := []int{}
			for j := 1; j < len(lineSplit); j++ {
				num, err := strconv.Atoi(lineSplit[j][:len(lineSplit[j])-1])
				if err != nil {
					logger.Fatal(err)
				}
				nums = append(nums, num)
			}
			posX, _ := strconv.Atoi(nameSplit[1][1:len(nameSplit[1])])
			posY, _ := strconv.Atoi(nameSplit[2][1:len(nameSplit[2])])

			n := node{
				name:  lineSplit[0],
				size:  nums[0],
				used:  nums[1],
				avail: nums[2],
				use:   nums[3],
				x:     posX,
				y:     posY,
			}
			nodes = append(nodes, n)
		}
	}

	// Print input nodes
	logger.Print("-------------------------- Nodes --------------------------")
	for _, n := range nodes {
		logger.Printf("%+v", n)
	}

	// Min steps to move data from node @ (maxX, 0) to (0,0)

	answer := 0
	return answer
}
