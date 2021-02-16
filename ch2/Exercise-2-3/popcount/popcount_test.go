package popcount

import (
	"testing"
)

var x uint64 = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(x)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(x)
	}
}

// output
// go test -bench=BenchmarkPopCount
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-3/popcount
// BenchmarkPopCount-8             1000000000               0.265 ns/op
// BenchmarkPopCountLoop-8         70633522                15.1 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-3/popcount        1.572s

// go test -bench=BenchmarkPopCountLoop
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-3/popcount
// BenchmarkPopCountLoop-8         74868979                15.1 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-3/popcount        1.316s
