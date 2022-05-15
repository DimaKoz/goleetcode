package easy

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	type Case struct {
		allowed string
		words   []string
		expect  int
	}
	cases := []Case{
		{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2},
		{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7},
		{"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4},
	}
	for index, item := range cases {
		if got := countConsistentStrings(item.allowed, item.words); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
