package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func tokenize(expr string) []string {
	var tokens []string
	var currToken string

	for _, chr := range expr {
		switch chr {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if len(currToken) > 0 {
				tokens = append(tokens, currToken)
				currToken = ""
			}
			tokens = append(tokens, string(chr))
		default:
			currToken += string(chr)
		}
	}

	if len(currToken) > 0 {
		tokens = append(tokens, currToken)
	}

	return tokens
}

func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	var out, op []string

	for _, token := range tokens {
		switch {
		case token == "(":
			op = append(op, token)
		case token == ")":
			for len(op) > 0 && op[len(op)-1] != "(" {
				out = append(out, op[len(op)-1])
				op = op[:len(op)-1]
			}
			if len(op) == 0 {
				return 0, errors.New("invalid expression")
			}
			op = op[:len(op)-1]
		case token == "+" || token == "-" || token == "*" || token == "/":
			for len(op) > 0 && (op[len(op)-1] == "*" || op[len(op)-1] == "/") {
				out = append(out, op[len(op)-1])
				op = op[:len(op)-1]
			}
			op = append(op, token)
		default:
			_, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression")
			}
			out = append(out, token)
		}
	}

	for len(op) > 0 {
		out = append(out, op[len(op)-1])
		op = op[:len(op)-1]
	}

	var stack []float64
	for _, token := range out {
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
