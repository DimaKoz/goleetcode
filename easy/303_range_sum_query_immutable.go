package easy

/*
https://leetcode.com/problems/range-sum-query-immutable/

Constraints:

1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
0 <= left <= right < nums.length
At most 10^4 calls will be made to sumRange.
*/

type NumArray struct {
	n *[]int
}

/*
Var2:
Runtime: 28 ms, faster than 84.65% of Go online submissions for Range Sum Query - Immutable.
Memory Usage: 8.4 MB, less than 92.98% of Go online submissions for Range Sum Query - Immutable.
*/
/*renamed Constructor -> ConstructorNumArrayVar2*/
func ConstructorNumArrayVar2(nums []int) NumArray {
	result := NumArray{&nums}
	for i := 1; i < len(*result.n); i++ {
		(*result.n)[i] += (*result.n)[i-1]
	}
	return result
}

/*renamed SumRange -> SumRangeVar2*/
func (this *NumArray) SumRangeVar2(left int, right int) int {
	if left == 0 {
		return (*this.n)[right]
	}
	return (*this.n)[right] - (*this.n)[left-1]
}

/*
Var1:
Runtime: 88 ms, faster than 19.30% of Go online submissions for Range Sum Query - Immutable.
Memory Usage: 9.5 MB, less than 23.68% of Go online submissions for Range Sum Query - Immutable.
*/
/*renamed Constructor -> ConstructorNumArrayVar1*/
func ConstructorNumArrayVar1(nums []int) NumArray {
	return NumArray{&nums}
}

/*renamed SumRange -> SumRangeVar1*/
func (this *NumArray) SumRangeVar1(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += (*this.n)[i]
	}
	return sum
}
