package easy

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Constraints:

The Linked List is not empty.
Number of nodes will not exceed 30.
Each node's value is either 0 or 1.
*/

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {

		//Runtime: 0 ms
		result = result*2 + head.Val
		//Runtime: 2 ms
		//result = result << 1 + head.Val

		head = head.Next
	}
	return result
}
