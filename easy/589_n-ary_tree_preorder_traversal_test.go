package easy

import "testing"

func TestPreorder(t *testing.T) {
	type Case struct {
		node   Node
		expect []int
	}
	n0 := Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}
	cases := []Case{
		{node: n0, expect: []int{1, 3, 5, 6, 2, 4}},
	}

	for _, item := range cases {
		if got := preorder(&item.node); len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != item.expect[i] {
					t.Errorf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, got[i], item.expect[i])
				}
			}
		}
	}
}
