package easy

import "testing"

func TestCheckTree(t *testing.T) {
	type Case struct {
		node   TreeNode
		expect bool
	}
	cases := []Case{
		{TreeNode{10, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}, true},
		{TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}}, false},
	}
	for index, item := range cases {
		if got := checkTree(&item.node); got != item.expect {
			t.Fatalf("Wrong with the case : %v, index: %d", item, index)
		}
	}
}
