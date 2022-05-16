package medium

import "testing"

func TestDiagonalSort(t *testing.T) {
	type Case struct {
		nums   [][]int
		expect [][]int
	}
	cases := []Case{
		{[][]int{
			{3, 3, 1, 1},
			{2, 2, 1, 2},
			{1, 1, 1, 2},
		}, [][]int{
			{1, 1, 1, 1},
			{1, 2, 2, 2},
			{1, 2, 3, 3},
		}},
		{[][]int{
			{11, 25, 66, 1, 69, 7},
			{23, 55, 17, 45, 15, 52},
			{75, 31, 36, 44, 58, 8},
			{22, 27, 33, 25, 68, 4},
			{84, 28, 14, 11, 5, 50},
		}, [][]int{
			{5, 17, 4, 1, 52, 7},
			{11, 11, 25, 45, 8, 69},
			{14, 23, 25, 44, 58, 15},
			{22, 27, 31, 36, 50, 66},
			{84, 28, 75, 33, 55, 68},
		},
		}}

	for _, item := range cases {
		got := diagonalSort(item.nums)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				for ii := 0; ii < len(item.expect[0]); ii++ {
					if got[i][ii] != item.expect[i][ii] {
						t.Fatalf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, item.nums[i], item.expect[i])
					}
				}

			}
		}
	}
}
