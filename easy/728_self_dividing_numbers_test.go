package easy

import "testing"

func TestSelfDividingNumbers(t *testing.T) {
	type Case struct {
		left   int
		right  int
		expect []int
	}
	cases := []Case{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{47, 85, []int{48, 55, 66, 77}},
	}

	for _, item := range cases {
		got := selfDividingNumbers(item.left, item.right)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
