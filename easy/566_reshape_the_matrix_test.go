package easy

import "testing"

func TestMatrixReshape(t *testing.T) {
	type Case struct {
		input  [][]int
		r      int
		c      int
		expect [][]int
	}
	cases := []Case{
		{input: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4, expect: [][]int{{1, 2}, {3, 4}}},
		{input: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4, expect: [][]int{{1, 2, 3, 4}}},
		{input: [][]int{{1, 2}, {3, 4}}, r: 4, c: 1, expect: [][]int{{1}, {2}, {3}, {4}}},
	}
	for _, item := range cases {
		got := matrixReshape(item.input, item.r, item.c)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i, row := range got {
				if len(row) != len(item.expect[i]) {
					t.Errorf("Wrong with the case : %v row: %v", item, row)
				}
				for j, a := range row {
					if a != item.expect[i][j] {
						t.Errorf("Wrong index[%d][%d] with the case : %v, got:%d, expect:%d", i, j, item, a, item.expect[i][j])
					}
				}
			}
		}
	}
}
