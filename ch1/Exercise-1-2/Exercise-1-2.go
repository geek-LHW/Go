// Exercise 1.2: Mo dif y the echo prog ram to print the index and
// value of each of its arguments,on e per line.
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
