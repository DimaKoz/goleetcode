package easy

import "testing"

func TestIsHappy(t *testing.T) {
	type Case struct {
		income int
		expect bool
	}
	cases := []Case{
		{income: 1, expect: true},
		{income: 19, expect: true},
		{income: 2, expect: false},
	}

	for _, item := range cases {
		if got := isHappy(item.income); got != item.expect {
			t.Errorf("Wrong with %v", item)
		}
	}
}
