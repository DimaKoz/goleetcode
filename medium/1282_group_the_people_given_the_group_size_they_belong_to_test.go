package medium

import "testing"

func TestGroupThePeople(t *testing.T) {
	type Case struct {
		groupSizes []int
		expect     [][]int
	}
	cases := []Case{
		{[]int{3, 3, 3, 3, 3, 1, 3}, [][]int{{5}, {0, 1, 2}, {3, 4, 6}}},
		{[]int{2, 1, 3, 3, 3, 2}, [][]int{{1}, {0, 5}, {2, 3, 4}}},
	}

	for _, item := range cases {
		got := groupThePeople(item.groupSizes)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				if len(got[i]) != len(item.expect[i]) {
					t.Errorf("Wrong row index[%d] with the case : %v, got:%d, expect:%d", i, item, got[i], item.expect[i])
				} else {
					for ii := 0; ii < len(item.expect[i]); ii++ {
						if got[i][ii] != item.expect[i][ii] {
							t.Fatalf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, got[i], item.expect[i])
						}
					}
				}
			}
		}
	}
}
