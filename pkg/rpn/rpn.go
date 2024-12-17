package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func tokenize(expr string) []string {
	var tokens []string
	var currentToken string

	for _, char := range expr {
		switch char {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if len(currentToken) > 0 {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		default:
			currentToken += string(char)
		}
	}

	if len(currentToken) > 0 {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	var output, operators []string

	for _, token := range tokens {
		switch {
		case token == "(":
			operators = append(operators, token)
		case token == ")":
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return 0, errors.New("invalid expression")
			}
			operators = operators[:len(operators)-1]
		case token == "+" || token == "-" || token == "*" || token == "/":
			for len(operators) > 0 && (operators[len(operators)-1] == "*" || operators[len(operators)-1] == "/") {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		default:
			_, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression")
			}
			output = append(output, token)
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	var stack []float64
	for _, token := range output {
		switch {
		case token == "+" || token == "-" || token == "*" || token == "/":
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, a/b)
			}
		default:
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token")
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}
