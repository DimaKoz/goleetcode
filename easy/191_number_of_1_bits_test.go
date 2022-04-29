package easy

import "testing"

func TestHammingWeight(t *testing.T) {
	var n uint32 = 00000000000000000000000010000000
	got := hammingWeight(n)
	if got != 1 {
		t.Fatal("wrong")
	}
}

func TestHammingWeightM(t *testing.T) {
	var n uint32 = 00000000000000000000000010000000
	got := hammingWeightM(n)
	if got != 1 {
		t.Fatal("wrong")
	}
}

func BenchmarkHammingWeightM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeightM(uint32(i))
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeight(uint32(i))
	}
}
