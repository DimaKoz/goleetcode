package medium

import "testing"

func TestExecuteInstructions(t *testing.T) {
	type Case struct {
		n        int
		startPos []int
		s        string
		expect   []int
	}
	cases := []Case{
		{3, []int{0, 1}, "RRDDLU", []int{1, 5, 4, 3, 1, 0}},
		{2, []int{1, 1}, "LURD", []int{4, 1, 0, 0}},
		{1, []int{0, 0}, "LURD", []int{0, 0, 0, 0}},
	}

	for _, item := range cases {
		got := executeInstructions(item.n, item.startPos, item.s)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				if got[i] != item.expect[i] {
					t.Errorf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, got[i], item.expect[i])
				}
			}
		}
	}
}
