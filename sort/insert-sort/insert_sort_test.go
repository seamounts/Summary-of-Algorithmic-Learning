package insert_sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	testCase := []struct {
		input  []int
		result []int
	}{
		{
			input: []int{
				1, 9, 5, 7, 2, 3, 5, 6, 8,
			},
			result: []int{
				1, 2, 3, 5, 5, 6, 7, 8, 9,
			},
		},
		{
			input: []int{
				1, -9, 50, 7, 20, 31, 5, 6, 8,
			},
			result: []int{
				-9, 1, 5, 6, 7, 8, 20, 31, 50,
			},
		},
		{
			input: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			result: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
		},
	}
	for _, tcase := range testCase {
		r := InsertSort(tcase.input)
		if len(tcase.result) != len(r) {
			t.Fatalf("want got [%v], actually: [%v]", tcase.result, r)
		}
		for i, _ := range r {
			if r[i] != tcase.result[i] {
				t.Fatalf("want got [%v], actually: [%v]", tcase.result, r)
			}
		}
		t.Logf("Sort completion: original seq: %v, now seq: %v", tcase.input, r)

	}

}
