package easy

/*
https://leetcode.com/problems/merge-two-sorted-lists/
Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
Runtime: 6 ms, faster than 21.47% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.5 MB, less than 46.24% of Go online submissions for Merge Two Sorted Lists.
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var temp = &ListNode{}
	var resultHead *ListNode = nil
	for list1 != nil || list2 != nil {
		var candidate *ListNode
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				candidate = list1
				list1 = list1.Next
			} else {
				candidate = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			candidate = list1
			list1 = list1.Next
		} else if list2 != nil {
			candidate = list2
			list2 = list2.Next
		} else {
			break
		}
		temp.Next = candidate
		temp = candidate
		if resultHead == nil {
			resultHead = candidate
		}
	}
	return resultHead
}
