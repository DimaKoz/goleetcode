package medium

/*
https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

Constraints:

n == boxes.length
1 <= n <= 2000
boxes[i] is either '0' or '1'.
*/

/*
Runtime: 3 ms, faster than 91.57% of Go online submissions for Minimum Number of Operations to Move All Balls to Each Box.
Memory Usage: 5 MB, less than 100.00% of Go online submissions for Minimum Number of Operations to Move All Balls to Each Box.
*/
func minOperations(boxes string) []int {
	operations := make([]int, len(boxes))
	var ballCounter, opsCounter, i int

	// -->
	for i = 0; i < len(operations); i++ {
		operations[i] = opsCounter
		if boxes[i] == '1' {
			ballCounter++
		}
		opsCounter += ballCounter
	}

	ballCounter, opsCounter = 0, 0

	//<--
	for i = len(operations) - 1; i >= 0 ;i-- {
		operations[i] += opsCounter
		if boxes[i] == '1' {
			ballCounter++
		}
		opsCounter += ballCounter
	}

	return operations
}

/*
Runtime: 72 ms, faster than 28.92% of Go online submissions for Minimum Number of Operations to Move All Balls to Each Box.
Memory Usage: 5.1 MB, less than 81.93% of Go online submissions for Minimum Number of Operations to Move All Balls to Each Box.
*/
func slowMinOperationsImpl(boxes string) []int {
	result := make([]int, len(boxes))

	absMinOperations := func(i int, j int) int {
		res := i - j
		if res < 0 {
			res = -res
		}
		return res
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			if j == i {
				continue
			}
			if boxes[j] == '0' {
				continue
			}
			abs := absMinOperations(i, j)
			result[i] += abs
		}
	}
	return result

}
