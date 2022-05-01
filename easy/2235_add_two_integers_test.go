package easy

import "testing"

func TestSum(t *testing.T) {
	type Case struct {
		num1   int
		num2   int
		expect int
	}
	cases := []Case{
		{12, 5, 17},
		{-10, 4, -6},
	}
	for index, item := range cases {
		if got := sum(item.num1, item.num2); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
