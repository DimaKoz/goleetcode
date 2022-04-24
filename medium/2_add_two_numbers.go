package medium

import (
	"math/big"
)

/*
https://leetcode.com/problems/add-two-numbers/

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersVar2(l1, l2)
}

/*
Runtime: 9 ms, faster than 74.67% of Go online submissions for Add Two Numbers.
Memory Usage: 6.7 MB, less than 5.00% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbersVar2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Slice := addTwoNumbersNodesToSlice(l1)
	l2Slice := addTwoNumbersNodesToSlice(l2)

	value1 := sliceToIntAddTwoNumbersReversed(l1Slice)
	value2 := sliceToIntAddTwoNumbersReversed(l2Slice)

	sum := value1.Add(value1, value2) // sum := value1 + value2

	result := splitToDigitsAddTwoNumbers(sum)

	var l3 = &ListNode{}
	var answer *ListNode = nil
	for _, item := range result {
		l3.Next = &ListNode{item, nil}
		if answer == nil {
			answer = l3.Next
		}
		l3 = l3.Next
	}
	if answer == nil {
		return &ListNode{}
	}
	return answer
}

/*
Runtime: 13 ms, faster than 50.78% of Go online submissions for Add Two Numbers.
Memory Usage: 6.5 MB, less than 5.00% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbersVar1(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Slice := addTwoNumbersNodesToSlice(l1)
	reverseIntAddTwoNumbers(l1Slice)
	l2Slice := addTwoNumbersNodesToSlice(l2)
	reverseIntAddTwoNumbers(l2Slice)

	value1 := sliceToIntAddTwoNumbers(l1Slice)

	value2 := sliceToIntAddTwoNumbers(l2Slice)

	sum := value1.Add(value1, value2) // sum := value1 + value2

	result := splitToDigitsAddTwoNumbers(sum)

	var l3 = &ListNode{}
	var answer *ListNode = nil
	for _, item := range result {
		l3.Next = &ListNode{item, nil}
		if answer == nil {
			answer = l3.Next
		}
		l3 = l3.Next
	}
	if answer == nil {
		return &ListNode{}
	}
	return answer
}

func addTwoNumbersNodesToSlice(l *ListNode) []int {
	nodes := make([]int, 0)
	head := l
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	return nodes
}

func sliceToIntAddTwoNumbersReversed(s []int) *big.Int {
	res := big.NewInt(0)
	op := big.NewInt(1)
	big10 := big.NewInt(10)
	for i := 0; i < len(s); i++ {
		//res += int64(s[i]) * op
		si := big.NewInt(int64(s[i])) // int64(s[i])
		siMulOp := si.Mul(si, op)     //int64(s[i]) * op
		res.Add(res, siMulOp)         //res += int64(s[i]) * op

		// op *= 10
		op.Mul(op, big10)

	}
	return res
}

func reverseIntAddTwoNumbers(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToDigitsAddTwoNumbers(n *big.Int) []int {
	var ret []int
	big10 := big.NewInt(10)
	big0 := big.NewInt(0)
	for n.Cmp(big0) != 0 {
		n10 := big.NewInt(0).Mod(n, big10) //n%10
		ret = append(ret /*n%10*/, int(n10.Int64()))
		n.Div(n, big10) //n /= 10

	}

	//we need a reversed result
	//so a slice is already nice
	//reverseIntAddTwoNumbers(ret)

	return ret
}

func sliceToIntAddTwoNumbers(s []int) *big.Int {
	res := big.NewInt(0)
	op := big.NewInt(1)
	for i := len(s) - 1; i >= 0; i-- {
		//res += int64(s[i]) * op
		si := big.NewInt(int64(s[i])) // int64(s[i])
		siMulOp := si.Mul(si, op)     //int64(s[i]) * op
		res.Add(res, siMulOp)         //res += int64(s[i]) * op

		// op *= 10
		op.Mul(op, big.NewInt(10))

	}
	return res
}
