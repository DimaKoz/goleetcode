package easy

import "testing"

func TestCountGoodTriplets(t *testing.T) {

	type Case struct {
		arr    []int
		a      int
		b      int
		c      int
		expect int
	}
	cases := []Case{
		{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
	}

	for _, item := range cases {
		got := countGoodTriplets(item.arr, item.a, item.b, item.c)
		if got != item.expect {
			t.Errorf(" Wrong with the case : expect %v, got:%v", item.expect, got)
		}
	}

}
