package easy

import "testing"

func TestMaxDepth(t *testing.T) {
	type Case struct {
		node   *TreeNode
		expect int
	}
	n0 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			9,
			nil,
			nil,
		},
		Right: &TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}

	cases := []Case{
		{node: n0, expect: 3},
	}

	for _, item := range cases {
		if got := maxDepth(item.node); got != item.expect {
			t.Errorf("Wrong with the case : %v", item)
		}
	}
}
