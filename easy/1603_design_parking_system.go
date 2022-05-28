package easy

/*
https://leetcode.com/problems/design-parking-system/

Constraints:

0 <= big, medium, small <= 1000
carType is 1, 2, or 3
At most 1000 calls will be made to addCar
*/

/*
Runtime: 42 ms, faster than 71.12% of Go online submissions for Design Parking System.
Memory Usage: 7.1 MB, less than 89.30% of Go online submissions for Design Parking System.
*/

type ParkingSystem struct {
	counter []int
}

/*Constructor -> ParkingSystemConstructor*/
func ParkingSystemConstructor(big int, medium int, small int) ParkingSystem {
	counter := []int{0, big, medium, small}
	system := ParkingSystem{}
	system.counter = counter
	return system
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.counter[carType] <= 0 {
		return false
	}
	this.counter[carType] -= 1
	return true
}
