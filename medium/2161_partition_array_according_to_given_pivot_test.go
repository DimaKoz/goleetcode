package medium

import "testing"

func TestPivotArray(t *testing.T) {
	type Case struct {
		nums   []int
		pivot  int
		expect []int
	}
	cases := []Case{
		{[]int{4, 0, 4, 5, -11}, 5, []int{4, 0, 4, -11, 5}},
		{[]int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{[]int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
	}

	for _, item := range cases {
		got := pivotArray(item.nums, item.pivot)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				if got[i] != item.expect[i] {
					t.Errorf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, item.nums[i], item.expect[i])
				}
			}
		}
	}
}
