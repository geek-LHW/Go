// 练习 1.11： 在fetchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。
// 如果一个网站没有回应，程序将采取怎样的行为？（Section8.9 描述了在这种情况下的应对机制）。
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
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
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

//output

// ./Exercise-1-11.exe https://www.google.com/ https://www.youtube.com/ https://www.qq.com/ https://www.taobao.com/

// Get "https://www.google.com/": x509: certificate signed by unknown authority
// 0.49s   102115  https://www.qq.com/
// 0.69s   120997  https://www.taobao.com/
// Get "https://www.youtube.com/": dial tcp 172.217.24.14:443: i/o timeout
// 30.00s elapsed
