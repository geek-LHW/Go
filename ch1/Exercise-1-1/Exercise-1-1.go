// Exercise 1.1: Mo dif y the echo prog ram to als o pr int os.Args[0],
// the name of the command that invo ked it
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
