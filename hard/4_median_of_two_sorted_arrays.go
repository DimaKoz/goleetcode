package hard

/*
https://leetcode.com/problems/median-of-two-sorted-arrays/
Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/

/*
Runtime: 20 ms, faster than 43.51% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.2 MB, less than 81.03% of Go online submissions for Median of Two Sorted Arrays.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// calculate the last number for search
	lenNum1 := len(nums1)
	lenNum2 := len(nums2)
	lenTotal := lenNum1 + lenNum2
	isOdd := lenTotal&1 == 1

	half := lenTotal / 2
	lastIndex := half

	// search for values of indexes
	var lastItem int
	lastIndexNums1, lastIndexNums2 := 0, 0
	var searchedValuesHalf, searchedValuesPreHalf, itemN1, itemN2 int
	for i := 0; i <= lastIndex; i++ {
		searchedValuesPreHalf = lastItem
		if lastIndexNums1 < lenNum1 && lastIndexNums2 < lenNum2 {
			itemN1 = nums1[lastIndexNums1]
			itemN2 = nums2[lastIndexNums2]
			if itemN1 < itemN2 {
				lastItem = itemN1
				lastIndexNums1++
			} else {
				lastItem = itemN2
				lastIndexNums2++
			}
		} else if lastIndexNums1 < lenNum1 { // nums2 done
			lastItem = nums1[lastIndexNums1]
			lastIndexNums1++
		} else if lastIndexNums2 < lenNum2 { // nums1 done
			lastItem = nums2[lastIndexNums2]
			lastIndexNums2++
		} else {
			break
		}

		if half > i {
			continue
		}
		searchedValuesHalf = lastItem

		break
	}

	if isOdd {
		return float64(searchedValuesHalf)
	} else {
		return float64(searchedValuesHalf+searchedValuesPreHalf) / 2.0
	}

}
