package easy

import "testing"

func TestCheckStraightLine(t *testing.T) {
	type Case struct {
		input  [][]int
		expect bool
	}
	cases := []Case{
		{input: [][]int{{0, 1}, {1, 3}, {-4, -7}, {5, 11}}, expect: true},
		{input: [][]int{{1, 2}, {2, 3}, {3, 5}}, expect: false},
		{input: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, expect: true},
		{input: [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, expect: false},
	}

	for index, item := range cases {
		if got := checkStraightLine(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
