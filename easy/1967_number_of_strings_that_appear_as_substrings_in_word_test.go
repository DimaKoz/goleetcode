package easy

import "testing"

func TestNumOfStrings(t *testing.T) {
	type Case struct {
		patterns []string
		word     string
		expect   int
	}
	cases := []Case{
		{[]string{"a", "abc", "bc", "d"}, "abc", 3},
		{[]string{"a", "b", "c"}, "aaaaabbbbb", 2},
		{[]string{"a", "a", "a"}, "ab", 3},
	}
	for index, item := range cases {
		if got := numOfStrings(item.patterns, item.word); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
