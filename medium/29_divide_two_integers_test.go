package medium

import "testing"

func TestDivide(t *testing.T) {
	type Case struct {
		dividend int
		divisor  int
		expect   int
	}
	cases := []Case{
		{-2147483648, -1, 2147483647},
		{10, 3, 3},
		{7, -3, -2},
	}

	for _, item := range cases {
		if got := divide(item.dividend, item.divisor); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
