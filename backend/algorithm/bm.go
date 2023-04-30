package algorithm

import (
	"math"
)

// BM is a function that implements the Boyer-Moore algorithm.
// @params text: string to be searched
// @params pattern: string to be searched for
// @return int:  idx of the first occurrence of pattern in text [0, len(text) - len(pattern))]
// 				-1 if pattern is not found in text,
//				-2 for invalid input (empty string, length of pattern > length of text, etc.)
func (a Algorithm) BM(text string, pattern string) int {
	// lowercase the strings and remove whitespace
	LowerAndTrim(&text)
	LowerAndTrim(&pattern)

	n := len(text)
	m := len(pattern)

	if n == 0 || m == 0 || m > n {
		// invalid input
		return -2
	}

	// build last occurence array
	l := buildLast(pattern)

	i := m - 1
	j := m - 1

	for i < n {
		if text[i] == pattern[j] {
			if j == 0 {
				// pattern found, return 1
				return i
			}
			i--
			j--
		} else {
			// shift pattern to the right
			i += m - int( math.Min( float64( j ), float64( 1 + l[text[i]] ) ) )
			j = m - 1
		}
	}

	// no match found
	return -1
}

// buildLast returns an array that stores the last occurrence of each character in the pattern.
// uses ASCII values to index the array.
// @params pattern: string to be searched for
// @return int[]: last occurence array of pattern
func buildLast(pattern string) []int {
	l := make([]int, 128) // last occurence array

	for i := 0; i < 128; i++ {
		l[i] = -1 // initialize all values to -1
	}

	for i := 0; i < len(pattern); i++ {
		l[pattern[i]] = i
	}

	return l
}