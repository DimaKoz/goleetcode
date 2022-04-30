package easy

/*
https://leetcode.com/problems/root-equals-sum-of-children/

Constraints:

The tree consists only of the root, its left child, and its right child.
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Root Equals Sum of Children.
Memory Usage: 2.1 MB, less than 43.37% of Go online submissions for Root Equals Sum of Children.
*/
func checkTree(root *TreeNode) bool {
	return root.Left.Val+root.Right.Val == root.Val
}
