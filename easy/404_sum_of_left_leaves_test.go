package easy

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
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
		{node: n0, expect: 24},
	}

	for _, item := range cases {
		if got := sumOfLeftLeaves(item.node); got != item.expect {
			t.Errorf("Wrong with the case : %v", item)
		}
	}
}
