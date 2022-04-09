package easy

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*Node, 0, 0)
	stack = append(stack, root)
	var node *Node
	var n, i int
	for len(stack) != 0 {

		n = len(stack) - 1 // Top element
		node = stack[n]
		stack = stack[:n] // Pop

		result = append(result, node.Val)

		for i = len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return result
}
