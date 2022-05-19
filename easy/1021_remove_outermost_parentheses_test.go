package easy

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	type Case struct {
		s      string
		expect string
	}
	cases := []Case{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
	}
	for index, item := range cases {
		if got := removeOuterParentheses(item.s); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
