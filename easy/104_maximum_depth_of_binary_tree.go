package easy

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Constraints:

The number of nodes in the tree is in the range [0, 10^4].
-100 <= Node.val <= 100
*/

func maxDepth(root *TreeNode) int {
	return recursionMaxDepth(root)
}

/*
Runtime: 3 ms, faster than 86.10% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.1 MB, less than 86.91% of Go online submissions for Maximum Depth of Binary Tree.
*/
func recursionMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getMax(recursionMaxDepth(root.Left), recursionMaxDepth(root.Right))
}

func getMax(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
