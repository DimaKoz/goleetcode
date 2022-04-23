package medium

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	type Case struct {
		l1     *ListNode
		l2     *ListNode
		expect *ListNode
	}
	cases := []Case{
		{
			l1:     &ListNode{0, nil},
			l2:     &ListNode{0, nil},
			expect: &ListNode{0, nil},
		},

		{
			l1:     &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:     &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			expect: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}

	for _, item := range cases {
		got := addTwoNumbers(item.l1, item.l2)
		checkResult(got, item.expect, t)
		got = addTwoNumbersVar1(item.l1, item.l2)
		checkResult(got, item.expect, t)
	}
}

func checkResult(got *ListNode, expect *ListNode, t *testing.T) {

	for got != nil && expect != nil {
		if got == nil || expect == nil {
			t.Errorf("different lenght of the case : %v, got:%v", expect, got)
			break
		}
		if got.Val != expect.Val {
			t.Errorf("Wrong with the case : %v, got:%v", expect, got)
		}
		got = got.Next
		expect = expect.Next
	}
}
