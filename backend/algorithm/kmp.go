package algorithm

// KMP is a function that implements the Knuth-Morris-Pratt algorithm.
// 	@params
// 		text: string to be searched
//		pattern: string to be searched for
// 	@return int:
//		idx of the first occurrence of pattern in text [0, len(text) - len(pattern))]
// 		-1 if pattern is not found in text,
//		-2 for invalid input (empty string, length of pattern > length of text, etc.)
func (a Algorithm) KMP(text string, pattern string) int {
	n := len(text)
	m := len(pattern)

	if n == 0 || m == 0 || m > n {
		// invalid input
		return -2
	}

	b := computeBorder(pattern)

	i := 0 // index of the text
	j := 0 // index of the pattern

	for i < n {
		if text[i] == pattern[j] {
			// if the characters match, then increment both indices
			if j == m-1 {
				// pattern found, return 1
				return i - len(pattern) + 1
			}
			i++
			j++
		} else if j > 0 {
			// if the characters don't match, then decrement the pattern index
			j = b[j-1]
		} else {
			// if the characters don't match and there is no previous border, then increment the text index
			i++
		}
	}
	// no match found
	return -1
}

// computeBorder is a helper function for KMP that computes the border array
// for the given pattern.
// @params pattern: string to be searched for
// @return int[]: border array of pattern
func computeBorder(pattern string) []int {
	b := make([]int, len(pattern)) // border array
	b[0] = 0                       // first element is always 0

	j := 0 // index of the border
	i := 1 // index of the element being computed

	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			// if the characters match, then the border is one more than the previous border
			b[i] = j + 1
			i++
			j++
		} else if j > 0 {
			// if the characters don't match, then the border is the border of the previous border
			j = b[j-1]
		} else {
			// if the characters don't match and there is no previous border, then the border is 0
			b[i] = 0
			i++
		}
	}
	return b
}
