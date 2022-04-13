package easy

import "testing"

func TestMoveZeroes(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{nums: []int{1, 0}, expect: []int{1, 0}},
		{nums: []int{1}, expect: []int{1}},
		{nums: []int{0, 0, 4}, expect: []int{4, 0, 0}},
		{nums: []int{0, 1, 0, 3, 12}, expect: []int{1, 3, 12, 0, 0}},
		{nums: []int{0}, expect: []int{0}},
	}

	for _, item := range cases {
		if moveZeroes(item.nums); len(item.nums) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.nums); i++ {
				if item.nums[i] != item.expect[i] {
					t.Errorf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, item.nums[i], item.expect[i])
				}
			}
		}
	}
}
