package easy

import "testing"

func TestCountKDifference(t *testing.T) {
	type Case struct {
		num    []int
		k      int
		expect int
	}
	cases := []Case{
		{[]int{1, 1, 2}, 1, 2},
		{[]int{9, 3, 1, 1, 4, 5, 4, 9, 5, 10}, 1, 8},
		{[]int{3, 2, 1, 5, 4}, 2, 3},
		{[]int{1, 3}, 3, 0},
		{[]int{1, 2, 2, 1}, 1, 4},
	}
	for index, item := range cases {
		if got := countKDifference(item.num, item.k); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d got%d", item, index, got)
		}
	}
}
