package easy

/*
https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

Constraints:

1 <= n <= 10^5
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
Memory Usage: 1.9 MB, less than 18.02% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
*/

func subtractProductAndSum(n int) int {
	productOfDigits := 1
	var sum, digit int
	for n != 0 {
		digit = n % 10
		productOfDigits *= digit
		sum += digit
		n /= 10
	}
	return productOfDigits - sum
}
