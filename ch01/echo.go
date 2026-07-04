package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	looper := 100000

	start := time.Now()
	for i := 0; i < looper; i++ {
		fmt.Fprintln(os.Stderr, strings.Join(os.Args[1:], " "))
	}
	diffJoin := time.Since(start)

	start = time.Now()
	for i := 0; i < looper; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Fprintln(os.Stderr, s)
	}
	diffPlus := time.Since(start)

	fmt.Println("time diff: ", diffJoin, diffPlus)
}
