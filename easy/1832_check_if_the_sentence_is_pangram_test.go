package easy

import "testing"

func TestCheckIfPangram(t *testing.T) {
	type Case struct {
		sentence string
		expect   bool
	}
	cases := []Case{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
	}
	for index, item := range cases {
		if got := checkIfPangram(item.sentence); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
