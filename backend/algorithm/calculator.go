package algorithm

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan persamaan matematika:")
	expr, _ := reader.ReadString('\n')
	expr = strings.TrimSpace(expr)
	// input := "-1 + 2 * 3 + ( -5  * 7)"
	preprocess := preprocessInput(expr)
	shunting := shuntingYard(preprocess)
	result, err := evaluatePostfix(shunting)
	if err != nil {
		fmt.Printf("Error       : '%s'\n", err)
	} else {
		fmt.Printf("Result      : '%f'\n", result)
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

func isDigit(token string) bool {
	_, err := strconv.Atoi(token)
	return err == nil
}

func preprocessInput(input string) string {
	input = strings.Replace(input, " ", "", -1)
	processed := strings.Builder{}

	for i, char := range input {
		if char == '-' && (i == 0 || input[i-1] == '(') {
			processed.WriteString(" 0 -")
		} else {
			processed.WriteRune(' ')
			processed.WriteRune(char)
		}
	}

	// Return trimmed string
	return strings.TrimSpace(processed.String())
}

func shuntingYard(input string) []string {
	output := []string{}
	operatorStack := []rune{}
	tokens := strings.Fields(input)

	for _, token := range tokens {
		tokenRune := []rune(token)
		if isDigit(token) {
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
		if isDigit(token) {
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
