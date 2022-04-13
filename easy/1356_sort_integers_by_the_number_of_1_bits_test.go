package easy

import "testing"

func TestSortByBits(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, expect: []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
		{nums: []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, expect: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
	}

	for _, item := range cases {
		if sortByBits(item.nums); len(item.nums) != len(item.expect) {
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
