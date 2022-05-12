package easy

import "testing"

func TestNumberOfMatches(t *testing.T) {
	type Case struct {
		n      int
		expect int
	}
	cases := []Case{
		{7, 6},
		{14, 13},
	}
	for index, item := range cases {
		if got := numberOfMatches(item.n); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
