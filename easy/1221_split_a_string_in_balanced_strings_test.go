package easy

import "testing"

func TestBalancedStringSplit(t *testing.T) {
	type Case struct {
		input  string
		expect int
	}
	cases := []Case{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
	}

	for index, item := range cases {
		if got := balancedStringSplit(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
