package easy

import "testing"

func TestCreateTargetArray(t *testing.T) {
	type Case struct {
		nums   []int
		index  []int
		expect []int
	}
	cases := []Case{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
		{[]int{1}, []int{0}, []int{1}},
	}

	for _, item := range cases {
		got := createTargetArray(item.nums, item.index)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
