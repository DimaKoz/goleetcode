package easy

import "testing"

func TestShuffle(t *testing.T) {
	type Case struct {
		nums   []int
		n      int
		expect []int
	}
	cases := []Case{
		{[]int{1, 2, 3, 4}, 2, []int{1, 3, 2, 4}},
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	}

	for _, item := range cases {
		got := shuffle(item.nums, item.n)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
