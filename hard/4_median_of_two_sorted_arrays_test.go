package hard

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	type Case struct {
		nums1  []int
		nums2  []int
		expect float64
	}
	cases := []Case{
		{[]int{1, 3}, []int{2}, 2.00000},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, item := range cases {
		if got := findMedianSortedArrays(item.nums1, item.nums2); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
