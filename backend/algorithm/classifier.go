// @NOT FINISHED
// package algorithm

// import (
// 	"regexp"
// )

// // Classify is a function that classifies the given text.
// // @params text: input string to be classified
// // @return int: sums of possible classifications
// //				 0: text is a question to be searched for in database
// //				 1: text is a math expression
// // 				 2: text is a date question
// // 				 4: text is a QA add request
// // 				 8: text is a QA delete request

// // @example return : text is a math expression and a date question
// //					return = 1 + 2 = 3
// func (a Algorithm) Classify(text string) int {
// 	res := 0

// 	// check if text is a math expression
// 	if IsMathExpression(text) {
// 		res += 1
// 	}

// 	// check if text is a date question
// 	if IsDateQuestion(text) {
// 		res += 2
// 	}

// 	// check if text is a QA add request
// 	if IsQAAddRequest(text) {
// 		res += 4
// 	}

// 	// check if text is a QA delete request
// 	if IsQADeleteRequest(text) {
// 		res += 8
// 	}

// 	return res
// }