package day_two

import (
	_ "embed"
	"testing"
)

//go:embed input
var partTwoInput []byte

func TestPartTwo(t *testing.T) {
	t.Log(partTwo(partTwoInput))
}

func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partTwo(partTwoInput)
	}
}
