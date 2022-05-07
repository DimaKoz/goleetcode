package easy

/*
https://leetcode.com/problems/shuffle-string/

Constraints:

s.length == indices.length == n
1 <= n <= 100
s consists of only lowercase English letters.
0 <= indices[i] < n
All values of indices are unique.
*/

/*
Runtime: 4 ms, faster than 77.01% of Go online submissions for Shuffle String.
Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Shuffle String.
*/

func restoreString(s string, indices []int) string {
	b := make([]byte, len(indices))
	for i := 0; i < len(indices); i++ {
		b[indices[i]] = s[i]
	}
	return string(b)
}
