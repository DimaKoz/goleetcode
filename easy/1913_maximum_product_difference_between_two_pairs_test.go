package easy

import "testing"

func TestMaxProductDifference(t *testing.T) {
	type Case struct {
		num    []int
		expect int
	}
	cases := []Case{
		{[]int{5, 6, 2, 7, 4}, 34},
		{[]int{4, 2, 5, 9, 7, 4, 8}, 64},
	}
	for index, item := range cases {
		if got := maxProductDifference(item.num); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d got%d", item, index, got)
		}
	}
}
