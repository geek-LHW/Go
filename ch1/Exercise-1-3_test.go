// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
//（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
package main

import (
	"strings"
	"testing"
)

func BenchmarkString2Join(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "JiangSu"}
		result := strings.Join(input, " ")
		if result != "Welcome To JiangSu" {
			b.Error("Unexcepted result:" + result)
		}
	}
}

func BenchmarkString2Plus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"Welcome", "To", "JiangSu"}
		var s, sep string
		for j := 0; j < len(input); j++ {
			s += sep + input[j]
			sep = " "
		}
		if s != "Welcome To JiangSu" {
			b.Error("Unexcepted result" + s)
		}
	}
}

// output

// go test -bench=BenchmarkString2Join
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go
// BenchmarkString2Join-8          16946758                60.2 ns/op
// PASS
// ok      github.com/geek-LHW/go  1.291s

// go test -bench=BenchmarkString2Plus
// goos: windows
// goarch: amd64
// pkg: github.com/geek-LHW/go
// BenchmarkString2Plus-8           9778515               111 ns/op
// PASS
// ok      github.com/geek-LHW/go  1.384s
