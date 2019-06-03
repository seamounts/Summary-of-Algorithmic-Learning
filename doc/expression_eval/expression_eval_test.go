package expression_eval

import (
	"testing"
)

func TestExpressionEval(t *testing.T) {
	testCase := []struct {
		input  string
		result float64
	}{
		{
			input:  "1+2+3-2",
			result: 4.0,
		},
		{
			input:  "1*2*3/2",
			result: 3.0,
		},
		{
			input:  "1*2*3/2+5/1*8+10-3",
			result: 50.0,
		},
	}
	for _, tcase := range testCase {
		r := ExpressionEval(tcase.input)
		if r != tcase.result {
			t.Fatalf("want got [%v], actually: [%v]", tcase.result, r)
		}
		t.Logf("ExpressionEval completion: input: %v, result: %v", tcase.input, r)
	}

}
