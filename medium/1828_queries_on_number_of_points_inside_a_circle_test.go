package medium

import "testing"

func TestCountPoints(t *testing.T) {
	type Case struct {
		points  [][]int
		queries [][]int
		expect  []int
	}
	cases := []Case{
		{[][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}, []int{3, 2, 2}},
		{[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}, []int{2, 3, 2, 4}},
	}

	for _, item := range cases {
		got := countPoints(item.points, item.queries)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				if got[i] != item.expect[i] {
					t.Errorf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, item.points[i], item.expect[i])
				}
			}
		}
	}
}
