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

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shift(x)
	}
}

// output
// go test -bench=BenchmarkPopCount
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-4/popcount
// BenchmarkPopCount-8             1000000000               0.269 ns/op
// BenchmarkPopCountShift-8        34331325                32.6 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-4/popcount        1.631s

// go test -bench=BenchmarkPopCountShift
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-4/popcount
// BenchmarkPopCountShift-8        30853087                32.8 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-4/popcount        1.224s
