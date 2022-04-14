package easy

import "testing"

func TestGetDecimalValue(t *testing.T) {
	type Case struct {
		input  *ListNode
		expect int
	}
	cases := []Case{
		{input: &ListNode{1, &ListNode{0, &ListNode{1, nil}}}, expect: 5},
		{input: &ListNode{0, nil}, expect: 0},
	}
	for index, item := range cases {
		if got := getDecimalValue(item.input); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
