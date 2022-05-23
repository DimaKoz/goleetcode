package easy

import "testing"

func TestMinimumSum(t *testing.T) {
	type Case struct {
		num    int
		expect int
	}
	cases := []Case{
		{2932, 52},
		{4009, 13},
	}
	for index, item := range cases {
		if got := minimumSum(item.num); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d got%d", item, index, got)
		}
	}
}
