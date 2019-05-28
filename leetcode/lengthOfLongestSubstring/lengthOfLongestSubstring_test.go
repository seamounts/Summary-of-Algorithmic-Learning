package lengthOfLongestSubstring

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	testCase := []struct { 
		input string
		result int
	}{
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"ab",
			2,
		},
		{
			"aab",
			2,
		},

	}
	for _, tcase := range testCase {
		r := LengthOfLongestSubstring(tcase.input)
		if r != tcase.result {
			t.Errorf("want got [%v], actually: [%v]", tcase.result, r)
		}
	}
}


func TestLengthOfLongestSubstring2(t *testing.T) {

	testCase := []struct { 
		input string
		result int
	}{
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"abcabcbb",
			3,
		},
		{
			"ab",
			2,
		},
		{
			"aab",
			2,
		},
		{
			"dvdf",
			3,
		},
		{
			"abba",
			2,
		},
		{
			"hkcpmprxxxqw",
			5,
		},
	}
	for _, tcase := range testCase {
		r := LengthOfLongestSubstring2(tcase.input)
		if r != tcase.result {
			t.Errorf("input [%v], want got [%v], actually: [%v]",tcase.input, tcase.result, r)
		}
	}
}

