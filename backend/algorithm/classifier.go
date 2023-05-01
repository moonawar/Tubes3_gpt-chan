package algorithm

import (
	"regexp"
)

// Classify is a function that classifies the given text.
// @params text: input string to be classified
// @return int: sums of possible classifications
//				 0: text is a question to be searched for in database
//				 1: text contains a candidate of math expression
// 				 2: text contains a date question
// 				 4: text contains a QA add request
// 				 8: text contains a QA delete request

// @example return : text is a math expression and a date question
//					return = 1 + 2 = 3
func (a Algorithm) Classify(text string) int {
	res := 0
	Lower(&text) // not case sensitive

	// check if text is a math expression
	if ContainsCandidateMathExp(text) {
		res += 1
	}

	// check if text is a date question
	if ContainsDate(text) {
		res += 2
	}

	// check if text is a QA add request
	if ContainsQAAddRequest(text) {
		res += 4
	}

	// check if text is a QA delete request
	if ContainsQADeleteRequest(text) {
		res += 8
	}

	return res
}

// ContainsCandidateMathExp is a function that checks if the given text contains a candidate math expression.
// @params text: input string to be checked
// @return bool: true if text contaions candidate math expression, false otherwise
//
// @note: candidate math expression does not necessarily correct math expression in terms of syntax
// @example: "1 + 2" is a candidate math expression, and "1 +* 2 " is still considered as a candidate math expression
//			 and thus both still return true. The syntax correctness is checked in the next step.
func ContainsCandidateMathExp(text string) {
	// remove whitespace so it's easier to check
	Trim(&text)

	// at least contains a binary expression
	regex_string := `\(?(\d+)\)?([\+\-\*\/\^]+)\(?(\d+)\)?`

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
//		  Tambah pertanyaan <pertanyaan> dengan jawaban <jawaban>
func ContainsQAAddRequest(text string) bool {
	regex_string := `[Tt]ambah(\s+)?[Pp]ertanyaan(\s+)?([\w\s]+)(\s+)?[Dd]engan(\s+)?[Jj]awaban(\s+)?([\w\s]+)`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}

// ContainsQADeleteRequest is a function that checks if the given text contains a QA delete request.
// @params text: input string to be checked
// @return bool: true if text contains a QA delete request, false otherwise
//
// @note: format: 
//		  Hapus pertanyaan <pertanyaan>
func ContainsQADeleteRequest(text string) bool {
	regex_string := `[Hh]apus(\s+)?[Pp]ertanyaan(\s+)?([\w\s]+)`

	re := regexp.MustCompile(regex_string)
	return re.MatchString(text)
}