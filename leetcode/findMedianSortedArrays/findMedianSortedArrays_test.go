package findMedianSortedArrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCase := []struct {
		input [2][]int
		result float64
	} {
		// {
		// 	[2][]int {
		// 		{1, 3},
		// 		{2},
		// 	},
		// 	2.0,
		// },
		// {
		// 	[2][]int {
		// 		{1, 2},
		// 		{3, 4},
		// 	},
		// 	2.5,
		// },
		// {
		// 	[2][]int {
		// 		{-2, -1},
		// 		{3},
		// 	},
		// 	-1,
		// },
		// {
		// 	[2][]int {
		// 		{},
		// 		{1},
		// 	},
		// 	1,
		// },
		{
			[2][]int {
				{4, 5, 6, 8, 9},
				{},
			},
			6,
		},
	}


	for _, tcase := range testCase {
		r := findMedianSortedArrays(tcase.input[0], tcase.input[1])
		if r != tcase.result {
			t.Errorf("want got [%v], actually: [%v]", tcase.result, r)
		}
	}
}