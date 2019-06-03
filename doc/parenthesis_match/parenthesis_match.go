package parenthesis_match

import (
	"github.com/seamounts/Summary-of-Algorithmic-Learning/stack"
)

func ParenthesisMatch(pm string) bool {
	PStack := stack.NewOrderStack()
	for i, _ := range pm {
		switch string(pm[i]) {
		case "{":
			fallthrough
		case "[":
			fallthrough
		case "(":
			PStack.Push(string(pm[i]))
		case "}":
			fallthrough
		case "]":
			fallthrough
		case ")":
			if PStack.Pop() != string(pm[i]) {
				return false
			}
		}
	}
	return PStack.Empty()
}
