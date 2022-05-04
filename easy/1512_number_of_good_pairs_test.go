package easy

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	type Case struct {
		input  []int
		expect int
	}
	cases := []Case{
		{input: []int{1, 2, 3, 1, 1, 3}, expect: 4},
		{input: []int{1, 1, 1, 1}, expect: 6},
		{input: []int{1, 2, 3}, expect: 0},
	}
	for index, item := range cases {
		if got := numIdenticalPairs(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
