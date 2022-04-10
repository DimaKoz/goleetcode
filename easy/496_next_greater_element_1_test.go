package easy

import "testing"

func TestNextGreaterElement(t *testing.T) {
	type Case struct {
		nums1  []int
		nums2  []int
		expect []int
	}

	cases := []Case{
		{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, expect: []int{3, -1}},
		{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, expect: []int{-1, 3, -1}},
	}

	for _, item := range cases {
		if got := nextGreaterElement(item.nums1, item.nums2); len(got) != len(item.expect) {
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
