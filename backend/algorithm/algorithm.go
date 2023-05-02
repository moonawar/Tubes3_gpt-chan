package algorithm

import (
	"strings"
)

type Algorithm struct{}

// New is a constructor for Algorithm.
func New() Algorithm {
	return Algorithm{}
}

// LowerAndTrim is a function that lowercases and removes all whitespace from a string.
func LowerAndTrim(s *string) {
	*s = strings.ToLower(*s)
	*s = strings.ReplaceAll(*s, " ", "")
}

// Lower is a function that lowercases a string.
func Lower(s *string) {
	*s = strings.ToLower(*s)
}

// Trim is a function that removes all whitespace from a string.
func Trim(s *string) {
	*s = strings.ReplaceAll(*s, " ", "")
}

// TrimFrontBack is a function that removes whitespace in the front and back only.
func TrimFrontBack(s *string) {
	*s = strings.TrimSpace(*s)
}

// HammingDistance is a function that calculates the hamming distance between two strings.
// @param s1: string 1
// @param s2: string 2
// @return: hamming distance, -1 if strings are not the same length
func HammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}

	var distance int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return distance
}
