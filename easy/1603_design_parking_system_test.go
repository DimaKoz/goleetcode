package easy

import "testing"

func TestParkingSystem(t *testing.T) {
	s := ParkingSystemConstructor(1, 1, 0)
	if got := s.AddCar(1); got != true {
		t.Errorf("Wrong with AddCar(1), got:%t", got)
	}
	if got := s.AddCar(2); got != true {
		t.Errorf("Wrong with AddCar(2), got:%t", got)
	}
	if got := s.AddCar(3); got != false {
		t.Errorf("Wrong with AddCar(3), got:%t", got)
	}
	if got := s.AddCar(1); got != false {
		t.Errorf("Wrong with AddCar(1), got:%t", got)
	}
}
