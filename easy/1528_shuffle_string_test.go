package easy

import "testing"

func TestRestoreString(t *testing.T) {
	type Case struct {
		input   string
		indices []int
		expect  string
	}
	cases := []Case{
		{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"abc", []int{0, 1, 2}, "abc"}}
	for index, item := range cases {
		if got := restoreString(item.input, item.indices); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
