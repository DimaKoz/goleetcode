package easy

import "testing"

func TestMaximumWealth(t *testing.T) {
	type Case struct {
		input  [][]int
		expect int
	}
	cases := []Case{
		{input: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, expect: 17},
		{input: [][]int{{1, 2, 3}, {3, 2, 1}}, expect: 6},
		{input: [][]int{{1, 5}, {7, 3}, {3, 5}}, expect: 10},
	}
	for index, item := range cases {
		if got := maximumWealth(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
