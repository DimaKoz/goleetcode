package easy

/*
https://leetcode.com/problems/reverse-string/

You must do this by modifying the input array in-place with O(1) extra memory.

Constraints:

1 <= s.length <= 10^5
s[i] is a printable ascii character.
*/

/*
Runtime: 29 ms, faster than 82.15% of Go online submissions for Reverse String.
Memory Usage: 6.6 MB, less than 93.36% of Go online submissions for Reverse String.
*/
func reverseString(s []byte) {
	lens := len(s)
	half := lens / 2
	lastIndex := lens - 1
	for i := 0; i < half; i++ {
		s[i], s[lastIndex] = s[lastIndex], s[i]
		lastIndex--
	}
}
