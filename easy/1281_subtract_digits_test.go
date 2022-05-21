package easy

import "testing"

func TestSubtractProductAndSum(t *testing.T) {
	type Case struct {
		input  int
		expect int
	}
	cases := []Case{
		{234, 15},
		{4421, 21},
	}

	for index, item := range cases {
		if got := subtractProductAndSum(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
