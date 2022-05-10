package easy

import "testing"

func TestCellsInRange(t *testing.T) {
	type Case struct {
		input  string
		expect []string
	}
	cases := []Case{
		{"A1:F1", []string{"A1", "B1", "C1", "D1", "E1", "F1"}},
		{"K1:L2", []string{"K1", "K2", "L1", "L2"}},
	}
	for index, item := range cases {

		if got := cellsInRange(item.input); len(got) != len(item.expect) {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != item.expect[i] {
					t.Fatalf("Wrong with the case : %v, index: %d at:%d", item, index, i)
				}
			}
		}
	}
}
