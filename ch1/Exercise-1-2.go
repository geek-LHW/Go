// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := ""
	for i, arg := range os.Args[1:] {
		s = strconv.Itoa(i) + " " + arg
		fmt.Println(s)
	}
}
