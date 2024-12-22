package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func makeToken(expr string) ([]string, error) {
	var tokens []string
	numToken := ""

	for _, chr := range expr {
		switch chr {
		case ' ':
			continue
		case '+', '-', '*', '/', '(', ')':
			if len(numToken) > 0 {
				tokens = append(tokens, numToken)
				numToken = ""
			}
			tokens = append(tokens, string(chr))
		default:
			if (chr < '0' || chr > '9') && chr != '.' {
				return nil, errors.New("invalid expression")
			}
			numToken += string(chr)
		}
	}

	if len(numToken) > 0 {
		tokens = append(tokens, numToken)
	}

	return tokens, nil
}

func rpnWrite(tokens []string) ([]string, error) {
	var output []string
	var temp []string

	var priority = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(temp) > 0 && priority[temp[len(temp)-1]] >= priority[token] {
				output = append(output, temp[len(temp)-1])
				temp = temp[:len(temp)-1]
			}

			temp = append(temp, token)

		case "(":
			temp = append(temp, token)

		case ")":
			for len(temp) > 0 && temp[len(temp)-1] != "(" {
				output = append(output, temp[len(temp)-1])
				temp = temp[:len(temp)-1]
			}

			if len(temp) == 0 {
				return nil, errors.New("invalid expression")
			}

			temp = temp[:len(temp)-1]
		default:
			output = append(output, token)
		}
	}

	for len(temp) > 0 {
		if temp[len(temp)-1] == "(" {
			return nil, errors.New("invalid expression")
		}
		output = append(output, temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}

	return output, nil
}

func Calc(expr string) (float64, error) {

	tokenList, err := makeToken(expr)

	if err != nil {
		return 0, err
	}

	rpnList, err := rpnWrite(tokenList)
	if err != nil {
		return 0, err
	}

	var stack []float64

	for _, token := range rpnList {
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
				return 0, fmt.Errorf("invalid expression")
			}
			stack = append(stack, num)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}
