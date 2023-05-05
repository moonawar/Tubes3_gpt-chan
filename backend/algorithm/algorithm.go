package algorithm

import (
	"math"
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
// @return: hamming distance in percentage (0 - 1), meaning
func (alg *Algorithm) HammingDistance(s1, s2 string) float64 {
	l := int(math.Min(float64(len(s1)), float64(len(s2))))

	distance := 0
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return 1 - (float64(distance) / float64(l))
}

func (a *Algorithm) LevenshteinDistance(s, t string) float64 {
	// Uses Pseudocode from https://en.wikipedia.org/wiki/Levenshtein_distance
	mat := make([][]int, len(s)+1)
	min_dist := int(math.Min(float64(len(s)), float64(len(t))))

	for i := 0; i < len(s)+1; i++ {
		mat[i] = make([]int, len(t)+1)
	}

	for i := 0; i < len(s)+1; i++ {
		for j := 0; j < len(t)+1; j++ {
			if j == 0 && i != 0 {
				mat[i][0] = i
			} else if i == 0 && j != 0 {
				mat[0][j] = j
			} else {
				mat[i][j] = 0
			}
		}
	}

	for j := 1; j < len(t)+1; j++ {
		for i := 1; i < len(s)+1; i++ {
			sub_cost := 1
			if s[i-1] == t[j-1] {
				sub_cost = 0
			}

			mat[i][j] = int(math.Min(
				float64(mat[i-1][j]+1),
				math.Min(
					float64(mat[i][j-1]+1),
					float64(mat[i-1][j-1]+sub_cost),
				),
			))
		}
	}

	return 1 - (float64(mat[len(s)][len(t)]) / float64(min_dist))
}
