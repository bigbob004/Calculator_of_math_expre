package parser_and_evaluater

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gammazero/deque"
	"os"
)

var priorities = map[byte]int{
	'(': 0,
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'~': 3,
}

const (
	LP int = iota
	RP
	NUMBER
	VAR
	OPERATION
)

type Token struct {
	code     string
	myType   int
	priority int
}

func IsDigit(sym byte) bool {
	if sym >= '0' && sym < '9' {
		return true
	}
	return false
}

func IsVariable(variable byte) bool {
	if variable < 'a' || variable > 'z' {
		return false
	}
	return true
}

func GetSringNumber(expr string, pos *int) (string, error) {
	number := []byte{}
	beg := *pos
	cntPoints := 0
	for ; beg < len(expr); beg++ {
		item := expr[beg]
		if item == '.' {
			cntPoints++
		}
		if cntPoints > 1 {
			return "", errors.New("Wrong Input")
		}
		if IsDigit(item) || item == '.' {
			number = append(number, item)
		} else {
			beg--
			break
		}
	}
	*pos = beg
	return string(number), nil
}

func SortedStation(tokens []Token) (out []Token, vars map[string]string, err error) {
	vars = map[string]string{}
	var stack deque.Deque
	for _, token := range tokens {
		if token.myType == VAR {
			//Возможно, нужно заполнение не нулями, а пустыми интерфейсами
			vars[token.code] = ""
			out = append(out, token)
		} else if token.myType == NUMBER {
			out = append(out, token)
			//Если символ является открывающей скобкой, помещаем его в стек.
		} else if token.myType == LP {
			stack.PushBack(token)
			//Если символ является закрывающей скобкой
		} else if token.myType == RP {
			for true {
				token = stack.PopBack().(Token)
				if token.myType == LP {
					break
				}
				if stack.Len() == 0 {
					return []Token{}, map[string]string{}, errors.New("Wrong input line")
				}
				out = append(out, token)
			}
		} else {
			//Если символ является бинарной операцией, тогда
			for stack.Len() != 0 && stack.Back().(Token).priority >= token.priority {
				out = append(out, stack.PopBack().(Token))
			}
			stack.PushBack(token)
		}
	}
	for stack.Len() != 0 {
		out = append(out, stack.PopBack().(Token))
	}
	if stack.Len() != 0 {
		return []Token{}, map[string]string{}, errors.New("Wrong input line")
	}
	return out, vars, nil
}

func InputVars(vars map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	for key := range vars {
		fmt.Printf("Введите значение переменной %s\n", key)
		val, _ := reader.ReadString('\n')
		vars[key] = val[:len(val)-1]
	}
}

func InsertValues(tokens []Token, vars map[string]string) {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.myType == VAR {
			tokens[i] = Token{code: vars[token.code], myType: NUMBER, priority: 0}
		}
	}
}

func Parsing(expr string) ([]Token, error) {
	out := []Token{}
	flag := 0
	var stack deque.Deque
	for i := 0; i < len(expr); i++ {
		s := expr[i]
		if IsVariable(s) {
			out = append(out, Token{code: string(s), myType: VAR, priority: 0})
			flag += 1
		} else if IsDigit(s) {
			if num, err := GetSringNumber(expr, &i); err != nil {
				return []Token{}, errors.New("Wrong Input")
			} else {
				out = append(out, Token{code: num, myType: NUMBER, priority: 0})
				flag += 1
			}
		} else if s == '(' {
			stack.PushBack('(')
			out = append(out, Token{code: string(s), myType: LP, priority: 0})
		} else if s == ')' {
			if stack.Len() == 0 {
				return []Token{}, errors.New("Wrong Input")
			}
			stack.PopBack()
			out = append(out, Token{code: string(s), myType: RP, priority: 0})
		} else if prior, ok := priorities[s]; ok {
			if s == '-' {
				if i == 0 {
					s = '~'
					prior = 3
					flag += 1
				} else if _, ok := priorities[expr[i-1]]; ok {
					s = '~'
					prior = 3
					flag += 1
				}
			}
			out = append(out, Token{code: string(s), myType: OPERATION, priority: prior})
			flag -= 1
		} else if s == ' ' {
			continue
		} else {
			return []Token{}, errors.New("Wrong Input")
		}
		if flag > 1 || flag < 0 {
			return []Token{}, errors.New("Wrong Input")
		}
	}
	if flag != 1 {
		return []Token{}, errors.New("Wrong Input")
	}
	if stack.Len() != 0 {
		return []Token{}, errors.New("Wrong Input")
	}
	return out, nil
}

func PrintExpr(expr []Token) {
	for _, token := range expr {
		fmt.Printf("%s ", token.code)
	}
	fmt.Print("\n")
}
