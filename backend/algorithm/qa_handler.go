package algorithm

import (
	"github.com/h2so5/goback/regexp"
)

// ExtractAddQuestion extract questions from a string in add request.
// @params text: string to be extracted
// @return question: question string, may be empty or multiple
func ExtractAddQuestions(text string) []string {
	regex_pattern := `(?<=[Tt]ambah(\s)[Pp]ertanyaan(\s))(.+)(?=[Dd]engan)`
	re := regexp.MustCompile(regex_pattern)
	text = RemoveExtraSpaces(text)

	res := re.FindAllString(text, -1)
	for i := 0; i < len(res); i++ {
		TrimFrontBack(&res[i])
	}
	return res
}

// ExtractAnswers extract answers from a string.
// @params text: string to be extracted
// @return answer: answer string, may be empty or multiple
func ExtractAnswers(text string) []string {
	regex_pattern := `(?<=[Dd]engan(\s)[Jj]awaban(\s))(.+)`
	re := regexp.MustCompile(regex_pattern)
	text = RemoveExtraSpaces(text)

	res := re.FindAllString(text, -1)
	for i := 0; i < len(res); i++ {
		TrimFrontBack(&res[i])
	}
	return res
}

// ExtractDeleteQuestions extract questions from a string in delete request.
// @params text: string to be extracted
// @return question: question string, may be empty or multiple
func ExtractDeleteQuestions(text string) []string {
	regex_pattern := `(?<=[Hh]apus(\s)[Pp]ertanyaan(\s))(.+)`
	re := regexp.MustCompile(regex_pattern)
	text = RemoveExtraSpaces(text)

	res := re.FindAllString(text, -1)
	for i := 0; i < len(res); i++ {
		TrimFrontBack(&res[i])
	}
	return res
}

// RemoveExtraSpaces removes extra spaces in a string.
// @params text: string to be processed
// @return string: string with extra spaces removed
// @example: "nama   aku    gpt" -> "nama aku gpt"
func RemoveExtraSpaces(text string) string {
	regex_pattern := `\s+`
	re := regexp.MustCompile(regex_pattern)

	return re.ReplaceAllString(text, " ")
}
