package easy

import "testing"

func TestReplaceDigits(t *testing.T) {
	type Case struct {
		s      string
		expect string
	}
	cases := []Case{
		{"a1c1e1", "abcdef"},
		{"a1b2c3d4e", "abbdcfdhe"},
	}
	for index, item := range cases {
		if got := replaceDigits(item.s); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
