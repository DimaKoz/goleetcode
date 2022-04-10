package easy

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	indexes := make(map[int]int, len(nums2))
	var len2 = len(nums2)
	for i := 0; i < len2; i++ {
		indexes[nums2[i]] = i
	}
	var add, idx int
	var found bool
	for i := 0; i < len(nums1); i++ {
		add = -1
		if idx, found = indexes[nums1[i]]; found {
			for ii := idx; ii < len2; ii++ {
				if nums1[i] < nums2[ii] {
					add = nums2[ii]
					break
				}
			}
		}
		result = append(result, add)
	}
	return result
}
