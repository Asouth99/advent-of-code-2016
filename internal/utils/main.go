package internal

import (
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Distance(v1, v2 []float64) float64 {
	var distance float64
	for i := range v1 {
		distance += math.Pow(v1[i]-v2[i], 2)
	}
	return math.Sqrt(distance) //, nil
}

func Combinations(n, m int) [][]int {
	var result [][]int
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				combination := make([]int, len(s))
				copy(combination, s)
				result = append(result, combination)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
	return result
}

func CombWithReplacement(n, m int) [][]int {
	var result [][]int
	s := make([]int, m)
	last := m - 1

	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				combination := make([]int, len(s))
				copy(combination, s)
				result = append(result, combination)
			} else {
				rc(i+1, j)
			}
		}
		return
	}
	rc(0, 0)
	return result
}

// Returns all permutations of a slice. eg: [1,2,3] returns [[1,2,3],[1,3,2],[2,3,1],[2,1,3],[3,1,2],[3,2,1]]
func GeneratePermutations[T comparable](arr []T) [][]T {
	var result [][]T
	var rc func(arr []T, n int)
	rc = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				rc(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	rc(arr, len(arr))
	return result
}

// Generates all ordered sequences of `n` positive integers that add up to exactly `sum`.
func Compositions(n int, sum int) [][]int {
	var results [][]int

	// Base edge cases
	if n <= 0 || sum < n {
		return results
	}

	var backtrack func(remainingElements int, remainingSum int, current []int)
	backtrack = func(remainingElements int, remainingSum int, current []int) {
		// Base case: If we are on the last number, it must be exactly the remaining sum
		if remainingElements == 1 {
			if remainingSum >= 0 {
				// Create a final copy of the composition to safely store it
				finalComposition := make([]int, len(current)+1)
				copy(finalComposition, current)
				finalComposition[len(current)] = remainingSum
				results = append(results, finalComposition)
			}
			return
		}

		// Try all possible positive integers for the current position
		// The maximum possible value leaves at least 1 for each remaining element
		maxPossible := remainingSum // - (remainingElements - 1)
		for i := 0; i < maxPossible+1; i++ {
			backtrack(remainingElements-1, remainingSum-i, append(current, i))
		}
	}

	// Start backtracking with an empty temporary slice
	backtrack(n, sum, make([]int, 0, n))
	return results
}

// This function only works for n > 0
// Returns a numbers factors
func Factorise(n int) []int {
	if n == 1 {
		return []int{1}
	}
	factors := []int{1, n} // Start with 1 and n
	maxFactor := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxFactor; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	return factors
}
