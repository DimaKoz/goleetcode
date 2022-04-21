package medium

import "testing"

func TestMyAtoi(t *testing.T) {
	type Case struct {
		expect int
		s      string
	}
	cases := []Case{
		{12345678, "  0000000000012345678"},
		{0, "  +  413"},
		{0, "00000-42a1234"},
		{0, "-+12"},
		{0, "+-12"},
		{-2147483648, "-91283472332"},
		{0, "words and 987"},
		{42, "42"},
		{-42, "   -42"},
		{4193, "4193 with words"},
	}

	for _, item := range cases {
		if got := myAtoi(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
