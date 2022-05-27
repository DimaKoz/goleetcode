package medium

import (
	"sort"
)

/*
https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/

Constraints:

groupSizes.length == n
1 <= n <= 500
1 <= groupSizes[i] <= n
*/

/*
Runtime: 3 ms, faster than 100.00% of Go online submissions for Group the People Given the Group Size They Belong To.
Memory Usage: 5.7 MB, less than 27.59% of Go online submissions for Group the People Given the Group Size They Belong To.
*/

func groupThePeople(groupSizes []int) [][]int {
	buckets := make(map[int][]int)
	found := false
	var bucket []int
	for i := 0; i < len(groupSizes); i++ {
		if bucket, found = buckets[groupSizes[i]]; found {
			bucket = append(bucket, i)
			buckets[groupSizes[i]] = bucket
		} else {
			bucket = make([]int, 0)
			bucket = append(bucket, i)
			buckets[groupSizes[i]] = bucket
		}
	}
	sort.Ints(groupSizes)
	result := make([][]int, 0)
	last := 0
	for _, key := range groupSizes {
		if last == key {
			continue
		}
		last = key
		bucket, _ = buckets[key]

		for len(bucket) > 0 {
			chunked := bucket[:key]
			result = append(result, chunked)
			bucket = bucket[len(chunked):]
		}
	}
	return result
}
