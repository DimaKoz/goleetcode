package easy

import "testing"

func TestFindCenter(t *testing.T) {
	type Case struct {
		edges  [][]int
		expect int
	}
	cases := []Case{
		{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
		{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, 1},
	}
	for index, item := range cases {
		if got := findCenter(item.edges); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
