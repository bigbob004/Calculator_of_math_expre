package parser_and_evaluater

import (
	"errors"
	"github.com/gammazero/deque"
	"strconv"
)

func binary_eval(rhs, lhs float64, operation string) (float64, error) {
	if operation == "+" {
		return rhs + lhs, nil
	}
	if operation == "-" {
		return lhs - rhs, nil
	}
	if operation == "*" {
		return rhs * lhs, nil
	}
	if rhs == 0 {
		return 0, errors.New("Zero division")
	}
	return lhs / rhs, nil
}

func unary_eval(item float64, operation string) (float64, error) {
	if operation == "~" {
		return item * -1, nil
	}
	return 0, errors.New("Something bad")
}

func Evaluate(expr []Token) (float64, error) {
	var stack deque.Deque
	for _, token := range expr {
		if token.myType == OPERATION {
			rhs := stack.PopBack().(float64)
			var res float64
			var err error
			if token.code == "~" {
				res, err = unary_eval(rhs, token.code)
			} else {
				lhs := stack.PopBack().(float64)
				res, err = binary_eval(rhs, lhs, token.code)
			}
			if err != nil {
				return 0, err
			}
			stack.PushBack(res)
		} else {
			num, _ := strconv.ParseFloat(token.code, 64)
			stack.PushBack(num)
		}
	}
	return stack.PopBack().(float64), nil
}
