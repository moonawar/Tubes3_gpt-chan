package algorithm

import (
	"regexp"
)

// Classify is a function that classifies the given text.
// @params text: input string to be classified
// @return int: sums of possible classifications
//				 0: text is a question to be searched for in database
//				 8: text contains a candidate of math expression
// 				 4: text contains a date question
// 				 2: text contains a QA add request
// 				 1: text contains a QA delete request

// @example return : text is a math expression and a date question
//
//	return = 1 + 2 = 3
func (a Algorithm) Classify(text string) int {
	res := 0

	// check if text contains candidate math expression
	if ContainsCandidateMathExp(text) {
		res += 8
	}

	// check if text contains date, date is cheked for edge case
	if ContainsDate(text) {
		res += 4
	}

	// check if text contains QA add request
	if ContainsQAAddRequest(text) {
		res += 2
	}

	// check if text contains QA delete request
	if ContainsQADeleteRequest(text) {
		res += 1
	}

	return res
}

// ContainsCandidateMathExp is a function that checks if the given text contains a candidate math expression.
// @params text: input string to be checked
// @return bool: true if text contaions candidate math expression, false otherwise
//
//	[]string: array of candidate math expression
//
// @note: candidate math expression does not necessarily correct math expression in terms of syntax
// @example: "1 + 2" is a candidate math expression, and "1 +* 2 " is still considered as a candidate math expression
//
//	and thus both still return true. The syntax correctness is checked in the next step.
func ContainsCandidateMathExp(text string) bool {
	// remove whitespace so it's easier to check
	Trim(&text)

	// recursively check for math expression
	regex_string := `(\d+|\(-?\d+\))(\s*[-+*/\s]*\s*(\d+|\(-?\d+\)))*`

	re := regexp.MustCompile(regex_string)
	all_exps := re.FindAllString(text, -1) // all candidate math expressions including date and unary expression
	pure_exps := []string{}                // remove date and unary expression

	for _, exp := range all_exps {
		if !ContainsDate(exp) && !IsUnaryMathExp(exp) {
			pure_exps = append(pure_exps, exp)
		}
	}

	return len(pure_exps) > 0
}

// IsUnaryMathExp is a function that checks if the given text is a unary math expression.
// @params text: input string to be checked
// @return bool: true if text is a unary math expression, false otherwise
func IsUnaryMathExp(text string) bool {
	// remove whitespace so it's easier to check
	Trim(&text)

	// recursively check for math expression
	regex_string := `^(\d+|\(-?\d+\))$`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}

// ContainsDate is a function that checks if the given text contains a date.
// @params text: input string to be checked
// @return bool: true if text contains a date, false otherwise
//
// @note: date format: dd/mm/yyyy
func ContainsDate(text string) bool {
	// remove whitespace so it's easier to check
	Trim(&text)

	regex_string := `(\d{1,2})[\/\-\\](\d{1,2})[\/\-\\](\d{4})`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}

// ContainsQAAddRequest is a function that checks if the given text contains a QA add request.
// @params text: input string to be checked
// @return bool: true if text contains a QA add request, false otherwise
//
// @note: format:
//
//	Tambah pertanyaan <pertanyaan> dengan jawaban <jawaban>
func ContainsQAAddRequest(text string) bool {
	regex_string := `[Tt]ambah(\s+)?[Pp]ertanyaan(\s+)?(.+)(\s+)?[Dd]engan(\s+)?[Jj]awaban(\s+)?(.+)`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}

// ContainsQADeleteRequest is a function that checks if the given text contains a QA delete request.
// @params text: input string to be checked
// @return bool: true if text contains a QA delete request, false otherwise
//
// @note: format:
//
//	Hapus pertanyaan <pertanyaan>
func ContainsQADeleteRequest(text string) bool {
	regex_string := `[Hh]apus(\s+)?[Pp]ertanyaan(\s+)?([\w\s]+)`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}

// ExtractMathExps is a function that extracts math expressions from the given text.
// @params text: input string to be checked
// @return []string: array of math expressions
func ExtractMathExps(text string) []string {
	regex_string := `(\d+|\(-?\d+\))(\s*[-+*/\s]*\s*(\d+|\(-?\d+\)))*`

	re := regexp.MustCompile(regex_string)
	return re.FindAllString(text, -1)
}
