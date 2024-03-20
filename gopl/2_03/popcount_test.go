package popcount

import (
	"math/rand"
	"testing"
)

var tests = []struct {
	input uint64
	want  int
}{
	{0x00000000, 0},
	{0x00000001, 1},
	{0x00010001, 2},
	{0x01010101, 4},
	{0x11111111, 8},
}

func TestPopCount(t *testing.T) {
	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%x) = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestPopCount1(t *testing.T) {
	for _, test := range tests {
		if got := PopCount1(test.input); got != test.want {
			t.Errorf("PopCount(%x) = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(rand.Uint64())
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount1(rand.Uint64())
	}
}
