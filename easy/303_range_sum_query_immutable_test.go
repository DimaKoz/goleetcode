package easy

import "testing"

func TestNumArray(t *testing.T) {
	type Case struct {
		nums   []int
		ranges [][]int
		expect []int
	}
	cases := []Case{
		{[]int{-2, 0, 3, -5, 2, -1}, [][]int{[]int{0, 2}, []int{2, 5}, []int{0, 5}}, []int{1, -1, -3}},
	}

	for _, item := range cases {
		num1Var1 := ConstructorNumArrayVar1(item.nums)
		for i := 0; i < len(item.ranges); i++ {
			if got := num1Var1.SumRangeVar1(item.ranges[i][0], item.ranges[i][1]); got != item.expect[i] {
				t.Errorf("SumRangeVar1 Wrong with the case : %v, got:%v", item, got)
			}
		}
	}

	for _, item := range cases {
		num1Var2 := ConstructorNumArrayVar2(item.nums)
		for i := 0; i < len(item.ranges); i++ {
			if got := num1Var2.SumRangeVar2(item.ranges[i][0], item.ranges[i][1]); got != item.expect[i] {
				t.Errorf("SumRangeVar2 Wrong with the case : %v, got:%v", item, got)
			}
		}
	}

}
