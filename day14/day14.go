package day14

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"strconv"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day14/input.txt"
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

// A hash is a key only if:
// - It contains three of the same character in a row, like 777. Only consider the first such triplet in a hash.
// - One of the next 1000 hashes in the stream contains that same character five times in a row, like 77777.
func containsTripleChar(hash string) (byte, bool) {
	for i := range len(hash) - 2 {
		if hash[i] == hash[i+1] && hash[i+1] == hash[i+2] {
			return hash[i], true
		}
	}
	return 0, false
}
func containsRepeatedChar(hash string, char byte, num int) bool {
	for i := range len(hash) - (num - 1) {
		if hash[i] != char {
			continue
		}
		allSame := true
		for j := i + 1; j < i+num; j++ {
			if hash[i] != hash[j] {
				allSame = false
				break
			}
		}
		if allSame {
			return true
		}
	}
	return false
}

func md5Hash(str string) [16]byte {
	return md5.Sum([]byte(str))
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	salt := string(f)
	keys := []int{}

	index := 1
	for len(keys) < 64 {
		indexStr := strconv.Itoa(index)
		hashBytes := md5Hash(salt + indexStr)
		hash := hex.EncodeToString(hashBytes[:])
		if char, ok := containsTripleChar(hash); ok {
			// logger.Printf("Index: %d, %s contains triple %c", index, hash, char)
			// Check if any of the next 1000 indexes contain 5 repeated chars
			contains5chars := false
			for index2 := index + 1; index2 <= index+1000; index2++ {
				indexStr2 := strconv.Itoa(index2)
				hashBytes2 := md5Hash(salt + indexStr2)
				hash2 := hex.EncodeToString(hashBytes2[:])
				if containsRepeatedChar(hash2, char, 5) {
					// logger.Printf("Index: %d, %s contains 5 %c", index2, hash, char)
					contains5chars = true
					break
				}
			}
			if contains5chars {
				keys = append(keys, index)
			}
		}
		index++
	}

	logger.Printf("Keys: %v", keys)

	answer := keys[63]
	return answer
}

func md5StretchHash(inputStr string, hashCache map[string][16]byte) [16]byte {
	if val, ok := hashCache[inputStr]; ok {
		return val
	}

	var hexBuf [32]byte
	var srcBuf [32]byte

	hash := md5.Sum([]byte(inputStr))
	hex.Encode(hexBuf[:], hash[:])

	for range 2016 {
		hash = md5.Sum(hexBuf[:])
		hex.Encode(srcBuf[:], hash[:])
		hexBuf = srcBuf
	}

	hashCache[inputStr] = hash
	return hash
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.ReadFile(inputFile)
	if err != nil {
		logger.Fatalf("error reading file: %v\n", err)
	}
	salt := string(f)
	keys := []int{}
	hashCache := map[string][16]byte{}
	index := 0
	for len(keys) < 64 {
		indexStr := strconv.Itoa(index)
		hashBytes := md5StretchHash(salt+indexStr, hashCache)
		// logger.Printf("idx: %d, %x", index, hashBytes)
		hash := hex.EncodeToString(hashBytes[:])
		if char, ok := containsTripleChar(hash); ok {
			// Check if any of the next 1000 indexes contain 5 repeated chars
			contains5chars := false
			for index2 := index + 1; index2 <= index+1000; index2++ {
				indexStr2 := strconv.Itoa(index2)
				hashBytes2 := md5StretchHash(salt+indexStr2, hashCache)
				hash2 := hex.EncodeToString(hashBytes2[:])
				if containsRepeatedChar(hash2, char, 5) {
					contains5chars = true
					break
				}
			}
			if contains5chars {
				keys = append(keys, index)
			}
		}
		index++
	}

	logger.Printf("Keys: %v", keys)

	answer := keys[len(keys)-1]
	return answer
}
