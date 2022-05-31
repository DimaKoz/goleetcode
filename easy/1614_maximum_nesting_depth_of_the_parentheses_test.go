package easy

import "testing"

func TestMaximumDepthParentheses(t *testing.T) {
	type Case struct {
		input  string
		expect int
	}
	cases := []Case{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
	}
	for index, item := range cases {
		if got := maxDepthParentheses(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
