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

func BenchmarkADV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ADV(x)
	}
}

// output
// go test -bench=BenchmarkPopCount
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-5/popcount
// BenchmarkPopCount-8     1000000000               0.274 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-5/popcount        0.478s

// go test -bench=BenchmarkADV
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go/ch2/Exercise-2-5/popcount
// BenchmarkADV-8          83281386                14.1 ns/op
// PASS
// ok      github.com/geek-LHW/go/ch2/Exercise-2-5/popcount        2.274s
