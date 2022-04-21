package medium

import "testing"

func TestCountAndSay(t *testing.T) {
	type Case struct {
		n      int
		expect string
	}
	cases := []Case{
		{2, "11"},
		{1, "1"},
		{4, "1211"},
	}

	for _, item := range cases {
		if got := countAndSay(item.n); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
