package easy

import "testing"

func TestSmallerNumbersThanCurrent(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{nums: []int{8, 1, 2, 2, 3}, expect: []int{4, 0, 1, 1, 3}},
		{nums: []int{6, 5, 4, 8}, expect: []int{2, 1, 0, 3}},
		{nums: []int{7, 7, 7, 7}, expect: []int{0, 0, 0, 0}},
	}

	for _, item := range cases {
		if smallerNumbersThanCurrent(item.nums); len(item.nums) != len(item.expect) {
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
