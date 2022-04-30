package easy

import "testing"

func TestRunningSum(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{nums: []int{1, 2, 3, 4}, expect: []int{1, 3, 6, 10}},
		{nums: []int{1, 1, 1, 1, 1}, expect: []int{1, 2, 3, 4, 5}},
		{nums: []int{3, 1, 2, 10, 1}, expect: []int{3, 4, 6, 16, 17}},
	}

	for _, item := range cases {
		got := runningSum(item.nums)
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
