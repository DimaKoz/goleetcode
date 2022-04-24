package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		s      string
		expect bool
	}
	cases := []Case{
		{"ab2a", false},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, item := range cases {
		if got := isPalindrome(item.s); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%t", item, got)
		}
	}
}
