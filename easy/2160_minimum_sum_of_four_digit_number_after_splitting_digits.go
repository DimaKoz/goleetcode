package easy

import "sort"

/*
https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/

Constraints:

1000 <= num <= 9999
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Sum of Four Digit Number After Splitting Digits.
Memory Usage: 1.9 MB, less than 81.88% of Go online submissions for Minimum Sum of Four Digit Number After Splitting Digits.
*/

func minimumSum(num int) int {
	digits := make([]int, 4)

	ind := 0
	for num > 0 {
		digits[ind] = num % 10
		num /= 10
		ind++
	}

	sort.Ints(digits)

	return digits[0]*10 + digits[2] /*new1*/ + digits[1]*10 + digits[3] /*new2*/
}
