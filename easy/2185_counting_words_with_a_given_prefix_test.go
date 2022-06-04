package easy

import "testing"

func TestPrefixCount(t *testing.T) {
	type Case struct {
		words  []string
		ref    string
		expect int
	}
	cases := []Case{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
	}
	for index, item := range cases {
		if got := prefixCount(item.words, item.ref); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
