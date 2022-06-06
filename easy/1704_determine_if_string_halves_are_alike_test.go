package easy

import "testing"

func TestHalvesAreAlike(t *testing.T) {
	type Case struct {
		s      string
		expect bool
	}
	cases := []Case{
		{"book", true},
		{"textbook", false},
	}

	for _, item := range cases {
		if got := halvesAreAlike(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
