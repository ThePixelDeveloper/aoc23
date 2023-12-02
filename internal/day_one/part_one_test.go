package day_one

import (
	_ "embed"
	"testing"
)

//go:embed input
var partOneInput []byte

func TestPartOne(t *testing.T) {
	t.Log(partOne(partOneInput))
}

func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partOne(partOneInput)
	}
}
