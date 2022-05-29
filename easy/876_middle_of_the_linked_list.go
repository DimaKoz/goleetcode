package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
https://leetcode.com/problems/middle-of-the-linked-list/

Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/

func middleNode(head *ListNode) *ListNode {
	return middleNodeSolution2(head)
}

/*
Runtime: 2 ms, faster than 37.95% of Go online submissions for Middle of the Linked List.
Memory Usage: 2 MB, less than 14.39% of Go online submissions for Middle of the Linked List.
*/
func middleNodeSolution1(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0, 100)
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}
	half := len(nodes) / 2
	return nodes[half]
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
Memory Usage: 2 MB, less than 14.39% of Go online submissions for Middle of the Linked List.
*/
func middleNodeSolution2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
