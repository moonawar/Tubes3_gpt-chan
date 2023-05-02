package algorithm

import (
	"strings"
)

type Algorithm struct {}

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