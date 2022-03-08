package main

import (
	src "Lab_1/pkg/parser_and_evaluater"
	"fmt"
)

func main() {
	//s := "(a/(s-d)+f-g*h+j*k)*(l*z+b/n-x/(c-v))"
	s := "-(-! * (4 + 3))"
	if expr, err := src.Parsing(s); err != nil {
		fmt.Print(err)
	} else if ok_expr, vars, err := src.SortedStation(expr); err != nil {
		fmt.Print(err)
	} else {
		src.PrintExpr(ok_expr)
		if len(vars) != 0 {
			src.InputVars(vars)
			fmt.Println(vars)
			src.InsertValues(ok_expr, vars)
		}
		if res, err := src.Evaluate(ok_expr); err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("res: %f", res)
		}
	}

}
