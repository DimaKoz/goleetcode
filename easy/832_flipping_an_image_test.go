package easy

import "testing"

func TestFlipAndInvertImage(t *testing.T) {
	type Case struct {
		nums   [][]int
		expect [][]int
	}
	cases := []Case{
		{[][]int{
			{1, 1, 0},
			{1, 0, 1},
			{0, 0, 0},
		}, [][]int{
			{1, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		}},
		{[][]int{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 0, 1, 0},
		}, [][]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
			{1, 0, 1, 0},
		},
		}}

	for _, item := range cases {
		got := flipAndInvertImage(item.nums)
		if len(got) != len(item.expect) {
			t.Errorf("Wrong with the case : %v", item)
		} else {
			for i := 0; i < len(item.expect); i++ {
				for ii := 0; ii < len(item.expect[0]); ii++ {
					if got[i][ii] != item.expect[i][ii] {
						t.Fatalf("Wrong index[%d] with the case : %v, got:%d, expect:%d", i, item, item.nums[i], item.expect[i])
					}
				}

			}
		}
	}
}
