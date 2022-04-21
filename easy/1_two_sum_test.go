package easy

import "testing"

func TestTwoSum(t *testing.T) {
	type Case struct {
		nums   []int
		target int
		expect []int // we can check a target
	}
	cases := []Case{
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, item := range cases {
		if got := twoSum(item.nums, item.target); len(got) != 2 ||
			(item.nums[got[0]]+item.nums[got[1]] != item.target) {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}

	}
}
