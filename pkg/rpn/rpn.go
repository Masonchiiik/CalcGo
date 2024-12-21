package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func tokenizer(expr string) ([]string, error) {
	var tokens []string
	currToken := ""

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
			if (chr < '0' || chr > '9') && chr != '.' {
				return nil, errors.New("invalid expression")
			}
			currToken += string(chr)
		}
	}

	if len(currToken) > 0 {
		tokens = append(tokens, currToken)
	}

	return tokens, nil
}

func rpnWrite(tokens []string) ([]string, error) {
	var out []string
	var tempStack []string

	var priority = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			for len(tempStack) > 0 && priority[tempStack[len(tempStack)-1]] >= priority[token] {
				out = append(out, tempStack[len(tempStack)-1])
				tempStack = tempStack[:len(tempStack)-1]
			}
			tempStack = append(tempStack, token)
		case "(":
			tempStack = append(tempStack, token)
		case ")":
			for len(tempStack) > 0 && tempStack[len(tempStack)-1] != "(" {
				out = append(out, tempStack[len(tempStack)-1])
				tempStack = tempStack[:len(tempStack)-1]
			}
			if len(tempStack) == 0 {
				return nil, errors.New("invalid expression")
			}
			tempStack = tempStack[:len(tempStack)-1]
		default:
			out = append(out, token)
		}
	}

	for len(tempStack) > 0 {
		if tempStack[len(tempStack)-1] == "(" {
			return nil, errors.New("invalid expression")
		}
		out = append(out, tempStack[len(tempStack)-1])
		tempStack = tempStack[:len(tempStack)-1]
	}

	return out, nil
}

func Calc(expr string) (float64, error) {

	tokenList, err := tokenizer(expr)

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
