package easy

import "testing"

func TestMergeTwoLists(t *testing.T) {
	type Case struct {
		list1  *ListNode
		list2  *ListNode
		expect []int
	}

	cases := []Case{
		{
			expect: []int{1, 2},
			list1:  &ListNode{1, nil},
			list2:  &ListNode{2, nil},
		},
		{
			expect: []int{1, 1, 2, 3, 4, 4},
			list1:  &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			list2:  &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		},
	}
	for _, item := range cases {
		got := mergeTwoLists(item.list1, item.list2)
		if got == nil {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		} else {
			index := -1
			for got != nil && index < len(item.expect)-1 {
				index++
				if got.Val != item.expect[index] {
					t.Errorf("Wrong with the case : %v, got:%v, item.expect[%d]:%d", item, got, index, item.expect[index])
				}
				got = got.Next
			}
		}
	}
}
