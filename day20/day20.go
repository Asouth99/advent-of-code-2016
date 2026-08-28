package day20

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day20/input.txt"
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
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// maxIp := 0
	// if strings.HasPrefix(inputFile, "example") {
	// 	maxIp = 9
	// } else {
	// 	maxIp = 4294967295
	// }

	blockedIpRanges := [][2]int{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Split(line, "-")
		min, _ := strconv.Atoi(lineSplit[0])
		max, _ := strconv.Atoi(lineSplit[1])

		blockedIpRanges = append(blockedIpRanges, [2]int{min, max})
	}

	// Loop through each IP and find lowest that is not blocked
	ip := 0
	for true {
		// Loop through each range and check if ip is blocked
		isBlocked := false
		for _, ipRange := range blockedIpRanges {
			if ip < ipRange[0] || ip > ipRange[1] {
				continue
			} else {
				ip = ipRange[1] // Move ip to the end of the range
				isBlocked = true
				break
			}
		}

		if !isBlocked {
			break
		}
		ip++
	}

	answer := ip
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	maxIp := 0
	if strings.HasPrefix(inputFile, "example") {
		maxIp = 9
	} else {
		maxIp = 4294967295
	}

	blockedIpRanges := [][2]int{}

	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Split(line, "-")
		min, _ := strconv.Atoi(lineSplit[0])
		max, _ := strconv.Atoi(lineSplit[1])

		blockedIpRanges = append(blockedIpRanges, [2]int{min, max})
	}

	// Loop through each IP and find ips that are allowed
	allowedIps := []int{}
	ip := 0
	for ip <= maxIp {
		// Loop through each range and check if ip is blocked
		isBlocked := false
		for _, ipRange := range blockedIpRanges {
			if ip < ipRange[0] || ip > ipRange[1] {
				continue
			} else {
				ip = ipRange[1] // Move ip to the end of the range
				isBlocked = true
				break
			}
		}

		if !isBlocked {
			allowedIps = append(allowedIps, ip)
		}
		ip++
	}

	answer := len(allowedIps)
	return answer
}
