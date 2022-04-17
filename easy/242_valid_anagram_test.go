package easy

import "testing"

func TestIsAnagram(t *testing.T) {
	type Case struct {
		s      string
		t      string
		expect bool
	}
	cases := []Case{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, item := range cases {
		if got := isAnagram(item.s, item.t); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%t", item, got)
		}
	}
}
