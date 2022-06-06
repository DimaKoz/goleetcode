package medium

import "testing"

func TestNumberOfBeams(t *testing.T) {
	type Case struct {
		bank   []string
		expect int
	}
	cases := []Case{
		{[]string{"011001", "000000", "010100", "001000"}, 8},
		{[]string{"000", "111", "000"}, 0},
	}

	for _, item := range cases {
		if got := numberOfBeams(item.bank); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
