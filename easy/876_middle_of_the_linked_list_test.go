package easy

import "testing"

func TestMiddleNode(t *testing.T) {
	type Case struct {
		head   *ListNode
		expect *ListNode
	}
	expect1 := &ListNode{3, &ListNode{4, &ListNode{5, nil}}}
	cases := []Case{
		{
			expect: expect1,
			head:   &ListNode{1, &ListNode{2, expect1}},
		},
	}
	for _, item := range cases {
		if got := middleNode(item.head); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
