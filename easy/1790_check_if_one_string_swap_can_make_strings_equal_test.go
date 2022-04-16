package easy

import "testing"

func TestAreAlmostEqual(t *testing.T) {
	type Case struct {
		s1     string
		s2     string
		expect bool
	}
	cases := []Case{
		{s1: "rbreb", s2: "brber1", expect: false},
		{s1: "rbreb", s2: "brber", expect: false},
		{s1: "bank", s2: "kanb", expect: true},
		{s1: "attack", s2: "defend", expect: false},
		{s1: "kelb", s2: "kelb", expect: true},
	}

	for _, item := range cases {
		if got := areAlmostEqual(item.s1, item.s2); got != item.expect {
			t.Errorf("Wrong with the case : %v", item)
		}
	}
}
