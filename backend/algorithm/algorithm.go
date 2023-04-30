package algorithm

import (
	"strings"
)

type Algorithm struct {}

// New is a constructor for Algorithm.
func NewAlgorithm() Algorithm {
	return Algorithm{}
}

func LowerAndTrim(s *string) {
	*s = strings.ToLower(*s)
	*s = strings.TrimSpace(*s)
}