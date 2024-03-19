package string_concat_test

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := os.Args[0]
		for i := 1; i < len(os.Args); i++ {
			s += " " + os.Args[i]
		}
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(os.Args, " ")
	}
}
