package easy

import "testing"

func TestArrayStringsAreEqual(t *testing.T) {
	type Case struct {
		word1  []string
		word2  []string
		expect bool
	}
	cases := []Case{
		{[]string{"ab", "c"}, []string{"a", "bcd"}, false},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
	}
	for index, item := range cases {
		if got := arrayStringsAreEqual(item.word1, item.word2); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d, %v", item, index, got)
		}
	}
}
