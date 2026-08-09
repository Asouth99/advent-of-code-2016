package day05

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day05/input.txt"
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

func SolvePart1(inputFile string, logger *log.Logger) string {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	logger.Print(string(f))
	input := string(f)

	password := ""

	num := -1
	for i := 0; i < 8; i++ {
		for true {
			num++
			str := input + strconv.Itoa(num)
			// if num%10000 == 0 {
			// 	logger.Printf("Checking %s", str)
			// }
			md5Sum := md5.Sum([]byte(str))
			hash := hex.EncodeToString(md5Sum[:])
			if hash[:5] == "00000" {
				logger.Printf("Found %d character in password at md5(%s) = %s", i, str, hash)
				password += string(hash[5])
				break
			}
		}
	}

	answer := password
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) string {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	logger.Print(string(f))
	input := string(f)

	// passwordMap := map[int]string{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: ""}
	passwordMap := map[int]string{}
	password := ""

	num := -1
	for len(passwordMap) < 8 {
		num++
		str := input + strconv.Itoa(num)
		md5Sum := md5.Sum([]byte(str))
		hash := hex.EncodeToString(md5Sum[:])

		if hash[:5] == "00000" {
			if hash[5] < '8' && hash[5] >= '0' {
				passwordIdx, err := strconv.Atoi(string(hash[5]))
				if err != nil {
					logger.Fatal(err)
				}
				if _, ok := passwordMap[passwordIdx]; !ok {
					passwordMap[passwordIdx] = string(hash[6])
					logger.Printf("Found %d character in password at md5(%s) = %s", passwordIdx, str, hash)
				}
			}
		}
	}

	for i := range 8 {
		password += passwordMap[i]
	}

	logger.Printf("Password: %s", password)

	answer := password
	return answer
}
