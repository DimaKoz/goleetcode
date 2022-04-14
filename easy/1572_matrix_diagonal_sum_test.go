package easy

import "testing"

func TestDiagonalSum(t *testing.T) {
	type Case struct {
		input  [][]int
		expect int
	}
	cases := []Case{
		{input: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, expect: 25},
		{input: [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, expect: 8},
		{input: [][]int{{5}}, expect: 5},
	}
	for index, item := range cases {
		if got := diagonalSum(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
