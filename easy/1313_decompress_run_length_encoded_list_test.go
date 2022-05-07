package easy

import "testing"

func TestDecompressRLElist(t *testing.T) {
	type Case struct {
		nums   []int
		expect []int
	}
	cases := []Case{
		{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
		{[]int{1, 1, 2, 3}, []int{1, 3, 3}},
	}

	for _, item := range cases {
		got := decompressRLElist(item.nums)
		if len(got) != len(item.expect) {
			t.Errorf(" Wrong length with the case : expect %v, got:%v", item.expect, got)
		} else {
			for i := 0; i < len(item.expect); i++ {
				if got[i] != item.expect[i] {
					t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
				}
			}
		}
	}

}
