package easy

import "testing"

func TestCountPoints(t *testing.T) {
	type Case struct {
		rings  string
		expect int
	}
	cases := []Case{
		{"B0B6G0R6R0R6G9", 1},
		{"B0R0G0R9R0B0G0", 1},
		{"G4", 0},
	}
	for index, item := range cases {
		if got := countPoints(item.rings); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
