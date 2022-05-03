package easy

import "testing"

func TestBuildArray(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}

	for _, item := range cases {
		got := buildArray(item.nums)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : %v, got:%v", item.nums, got)
			}
		}
	}

}
