package easy

import "testing"

func TestOrderedStream(t *testing.T) {

	ordS := ConstructorOrderedStream(5)
	if got := ordS.InsertOrderedStream(3, "ccccc"); len(got) != 0 {
		t.Fatalf("Wrong with the case : %v", "ccccc")
	}
	expect := []string{"aaaaa"}
	if got := ordS.InsertOrderedStream(1, "aaaaa"); len(got) != len(expect) {
		t.Fatalf("Wrong got : %v, expect: %v", got, expect)
	} else {
		for i := 0; i < len(expect); i++ {
			if expect[i] != got[i] {
				t.Fatalf("Wrong got : %v, expect: %v", got[i], expect[i])
			}
		}
	}
	expect = []string{"bbbbb", "ccccc"}
	if got := ordS.InsertOrderedStream(2, "bbbbb"); len(got) != len(expect) {
		t.Fatalf("Wrong got : %v, expect: %v", got, expect)
	} else {
		for i := 0; i < len(expect); i++ {
			if expect[i] != got[i] {
				t.Fatalf("Wrong got : %v, expect: %v", got[i], expect[i])
			}
		}
	}
	expect = []string{}
	if got := ordS.InsertOrderedStream(5, "eeeee"); len(got) != len(expect) {
		t.Fatalf("Wrong got : %v, expect: %v", got, expect)
	}
	expect = []string{"ddddd", "eeeee"}
	if got := ordS.InsertOrderedStream(4, "ddddd"); len(got) != len(expect) {
		t.Fatalf("Wrong got : %v, expect: %v", got, expect)
	} else {
		for i := 0; i < len(expect); i++ {
			if expect[i] != got[i] {
				t.Fatalf("Wrong got : %v, expect: %v", got[i], expect[i])
			}
		}
	}

}
