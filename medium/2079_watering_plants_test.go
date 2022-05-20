package medium

import "testing"

func TestWateringPlants(t *testing.T) {
	type Case struct {
		plants   []int
		capacity int
		expect   int
	}
	cases := []Case{
		{[]int{2, 2, 3, 3}, 5, 14},
		{[]int{1, 1, 1, 4, 2, 3}, 4, 30},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 8, 49},
	}

	for _, item := range cases {
		if got := wateringPlants(item.plants, item.capacity); got != item.expect {
			t.Errorf("Wrong with the case : %v, got:%v", item, got)
		}
	}
}
