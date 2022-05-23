package easy

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	type Case struct {
		income []int
		expect bool
	}
	cases := []Case{
		{income: []int{-6, -2, -4}, expect: true},
		{income: []int{3, 5, 1}, expect: true},
		{income: []int{1, 2, 4}, expect: false},
		{income: []int{1, 2}, expect: true},
		{income: []int{1, 1}, expect: true},
	}

	for _, item := range cases {
		if got := canMakeArithmeticProgression(item.income); got != item.expect {
			t.Fatalf("Wrong with %v", item)
		}
	}
}
