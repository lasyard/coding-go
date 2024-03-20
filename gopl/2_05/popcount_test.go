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

func TestPopCount3(t *testing.T) {
	for _, test := range tests {
		if got := PopCount3(test.input); got != test.want {
			t.Errorf("PopCount(%x) = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount3(rand.Uint64())
	}
}
