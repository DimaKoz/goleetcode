package easy

import "testing"

func TestSumOddLengthSubarrays(t *testing.T) {
	type Case struct {
		input  []int
		expect int
	}
	cases := []Case{
		{input: []int{1, 4, 2, 5, 3}, expect: 58},
		{input: []int{1, 2}, expect: 3},
	}

	for index, item := range cases {
		if got := sumOddLengthSubarrays(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
