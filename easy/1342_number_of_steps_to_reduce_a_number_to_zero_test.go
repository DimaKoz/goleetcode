package easy

import "testing"

func TestNumberOfSteps(t *testing.T) {
	type Case struct {
		input  int
		expect int
	}
	cases := []Case{
		{0, 0},
		{14, 6},
		{8, 4},
		{123, 12},
	}

	for index, item := range cases {
		if got := numberOfSteps(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
