package easy

import "testing"

func TestCountPairs(t *testing.T) {
	type Case struct {
		nums   []int
		k      int
		expect int
	}
	cases := []Case{
		{[]int{3, 1, 2, 2, 2, 1, 3}, 2, 4},
		{[]int{1, 2, 3, 4}, 1, 0},
	}
	for index, item := range cases {
		if got := countPairs(item.nums, item.k); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d got%d", item, index, got)
		}
	}
}
