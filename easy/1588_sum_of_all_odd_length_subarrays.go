package easy

func sumOddLengthSubarrays(arr []int) int {
	lArr := len(arr)
	sum := 0
	var subA []int
	for odd := 1; odd <= lArr; odd += 2 {
		for i := 0; odd+i <= lArr; i++ {
			subA = arr[i : odd+i]
			for _, item := range subA {
				sum += item
			}
		}
	}
	return sum
}
