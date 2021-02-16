// Exercise 1.10: Findaweb sit e that pro duces a large amount of dat a. Invest igate caching by
// running fetchall twice in succession to see whether the rep orted time changes much. Do
// you get the same content each time? Modif y fetchall to print its out put to a file so it can be
// examined
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		fmt.Printf("start fetch %v\n", url)
		go fetch(url, ch) // start a goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// ./Exercise-1-10.exe http://www.baidu.com  http://taobao.com
//output
// 0.24s   297681  http://www.baidu.com
// 0.29s   297647  http://www.baidu.com
// 0.73s   120997  http://taobao.com
// 0.80s   120997  http://taobao.com
// 0.81s elapsed
