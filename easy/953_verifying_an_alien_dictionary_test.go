package easy

import "testing"

func TestIsAlienSorted(t *testing.T) {
	type Case struct {
		words  []string
		order  string
		expect bool
	}
	cases := []Case{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{[]string{"app", "apple"}, "abcdefghijklmnopqrstuvwxyz", true},
	}

	for _, item := range cases {
		if got := isAlienSorted(item.words, item.order); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
