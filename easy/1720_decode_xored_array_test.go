package easy

import "testing"

func TestDecode(t *testing.T) {
	type Case struct {
		encoded []int
		first   int
		expect  []int
	}
	cases := []Case{
		{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
		{[]int{6, 2, 7, 3}, 4, []int{4, 2, 0, 7, 4}},
	}

	for _, item := range cases {
		got := decode(item.encoded, item.first)
		for i := 0; i < len(item.expect); i++ {
			if got[i] != item.expect[i] {
				t.Errorf(" Wrong with the case : expect %v, got:%v at:%d", item.expect, got, i)
			}
		}
	}

}
