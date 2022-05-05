package easy

import "testing"

func TestKidsWithCandies(t *testing.T) {
	type Case struct {
		candies      []int
		extraCandies int
		expect       []bool
	}
	cases := []Case{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, item := range cases {
		got := kidsWithCandies(item.candies, item.extraCandies)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
