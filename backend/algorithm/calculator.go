package algorithm

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var tokenPatterns = [...]string{`^\(`, `^\)`, `^\+`, `^-`, `^\*`, `^/`, `^\^`, `^(\d*\.)?\d+`}

func (a *Algorithm) SolveMath(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	// input := "-1 + 2 * 3 + ( -5  * 7)"
	preprocess := preprocessInput(expr)
	shunting := shuntingYard(preprocess)
	result, err := evaluatePostfix(shunting)
	if err != nil {
		return 0, err
	} else {
		return result, nil
	}
}

type Operator struct {
	precedence        int
	fn                func(float64, float64) (float64, error)
	associativityLeft bool
}

var operatorsMap = map[rune]Operator{
	'+': {1, func(a, b float64) (float64, error) { return a + b, nil }, true},
	'-': {1, func(a, b float64) (float64, error) { return a - b, nil }, true},
	'*': {2, func(a, b float64) (float64, error) { return a * b, nil }, true},
	'/': {2, func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("invalid division by zero")
		}
		return a / b, nil
	}, true},
	'^': {3, func(a, b float64) (float64, error) { return math.Pow(a, b), nil }, false},
}

func isOperator(token rune) bool {
	_, exists := operatorsMap[token]
	return exists
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}

func preprocessInput(input string) string {
	input = strings.Replace(input, " ", "", -1)
	processed := strings.Builder{}

	for i, char := range input {
		if char == '-' && (i == 0 || input[i-1] == '(') {
			processed.WriteString(" 0 -")
		} else {
			processed.WriteRune(char)
		}
	}

	// Return trimmed string
	return strings.TrimSpace(processed.String())
}

func tokenize(input string) []string {
	var tokens []string
	regexes := make([]*regexp.Regexp, len(tokenPatterns))
	for i, pattern := range tokenPatterns {
		regexes[i] = regexp.MustCompile(pattern)
	}
	for len(input) > 0 {
		input = strings.TrimSpace(input)
		found := false
	inner:
		for _, re := range regexes {
			token := re.FindString(input)
			if len(token) > 0 {
				tokens = append(tokens, token)
				input = input[len(token):]
				found = true
				break inner
			}
		}
		if !found {
			return []string{}
		}
	}
	return tokens
}

func shuntingYard(input string) []string {
	output := []string{}
	operatorStack := []rune{}
	tokens := tokenize(input)

	for _, token := range tokens {
		tokenRune := []rune(token)
		if isNumber(token) {
			output = append(output, token)
		} else if isOperator(tokenRune[0]) {
			op1 := []rune(token)[0]
			for len(operatorStack) > 0 {
				op2 := operatorStack[len(operatorStack)-1]
				if isOperator(op2) && ((operatorsMap[op1].associativityLeft && operatorsMap[op1].precedence <= operatorsMap[op2].precedence) || (operatorsMap[op1].precedence < operatorsMap[op2].precedence)) {
					output = append(output, string(op2))
					operatorStack = operatorStack[:len(operatorStack)-1]
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, tokenRune[0])
		} else if tokenRune[0] == '(' {
			operatorStack = append(operatorStack, tokenRune[0])
		} else if tokenRune[0] == ')' {
			for {
				op := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if op == '(' {
					break
				}
				output = append(output, string(op))
			}
		}
	}

	for len(operatorStack) > 0 {
		output = append(output, string(operatorStack[len(operatorStack)-1]))
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return output
}

func evaluatePostfix(postfix []string) (float64, error) {
	stack := []float64{}

	for _, token := range postfix {
		tokenRune := []rune(token)
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(tokenRune[0]) {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result, error := operatorsMap[tokenRune[0]].fn(a, b)
			if error != nil {
				return 0, error
			}
			stack = append(stack, result)
		}
	}

	return stack[0], nil
}
