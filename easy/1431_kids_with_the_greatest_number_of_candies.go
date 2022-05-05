package easy

/*
https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

Constraints:

n == candies.length
2 <= n <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50

*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Kids With the Greatest Number of Candies.
Memory Usage: 2.4 MB, less than 69.31% of Go online submissions for Kids With the Greatest Number of Candies.
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxBeforeExtra := 0
	for i := 0; i < len(candies); i++ {
		if maxBeforeExtra < candies[i] {
			maxBeforeExtra = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		result[i] = maxBeforeExtra <= candies[i]+extraCandies
	}
	return result
}
