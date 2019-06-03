package expression_eval

import (
	"strconv"

	"github.com/seamounts/Summary-of-Algorithmic-Learning/stack"
)

func ExpressionEval(exp string) float64 {
	operatorStack := stack.NewOrderStack()
	numberStack := stack.NewOrderStack()
	var result float64
	var index int
	for i, _ := range exp {
		switch string(exp[i]) {
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			numberStack.Push(exp[index:i])
			if operatorStack.Get() == "*" || operatorStack.Get() == "/" {
				num1, err := strconv.ParseFloat(numberStack.Pop(), 64)
				if err != nil {
					panic(err)
				}
				num2, err := strconv.ParseFloat(numberStack.Pop(), 64)
				if err != nil {
					panic(err)
				}
				result = getResult(num2, num1, operatorStack.Pop())
				numberStack.Push(strconv.FormatFloat(result, 'f', 6, 64))
			}
			operatorStack.Push(string(exp[i]))
			index = i + 1
		}
	}
	if index != len(exp) {
		numberStack.Push(exp[index:len(exp)])
	}

	for !operatorStack.Empty() {
		op := operatorStack.Pop()
		num1, err := strconv.ParseFloat(numberStack.Pop(), 64)

		if err != nil {
			panic(err)
		}
		num2, err := strconv.ParseFloat(numberStack.Pop(), 64)
		if err != nil {
			panic(err)
		}

		result = getResult(num2, num1, op)
		numberStack.Push(strconv.FormatFloat(result, 'f', 6, 64))
	}
	return result
}

func getResult(num1, num2 float64, operator string) float64 {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0.0
}
