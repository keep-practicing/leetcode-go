package allanagramsinastring

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type arg struct {
		s string
		p string
	}
	cases := []arg{
		arg{
			s: "cbaebabacd",
			p: "abc",
		},
		arg{
			s: "abab",
			p: "ab",
		},
		arg{
			p: "abab",
			s: "ab",
		},
	}

	expected := [][]int{
		{0, 6},
		{0, 1, 2},
		{},
	}

	for index, args := range cases {
		if res := findAnagrams(args.s, args.p); !reflect.DeepEqual(res, expected[index]) {
			t.Errorf("expected %v, got %v", expected[index], res)
		}
	}
}
