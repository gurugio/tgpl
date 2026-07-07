package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go ex10Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func ex10Fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

/*
user@AL02279337 ch01 % go run fetchall.go https://golang.org http://gopl.io https://godoc.org
1.02   34115 https://godoc.org
1.09   64185 https://golang.org
1.19    4154 http://gopl.io
1.19s elapsed
user@AL02279337 ch01 % go run fetchall.go https://golang.org http://gopl.io https://godoc.org
0.71   34115 https://godoc.org
0.86    4154 http://gopl.io
1.00   64185 https://golang.org
1.00s elapsed
user@AL02279337 ch01 % go run fetchall.go https://golang.org http://gopl.io https://godoc.org
0.89    4154 http://gopl.io
0.96   34115 https://godoc.org
0.98   64185 https://golang.org
0.98s elapsed
user@AL02279337 ch01 % go run fetchall.go https://golang.org http://gopl.io https://godoc.org
0.68   34115 https://godoc.org
0.85    4154 http://gopl.io
0.89   64185 https://golang.org
0.89s elapsed
user@AL02279337 ch01 % go run fetchall.go https://golang.org http://gopl.io https://godoc.org
0.73   34115 https://godoc.org
0.84    4154 http://gopl.io
0.85   64185 https://golang.org
0.85s elapsed
*/
