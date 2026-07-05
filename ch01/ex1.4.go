package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Printf("Arguments with File names are mandatory")
		return
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesFromFile(f, counts)
			f.Close()
		}
	}
	for line, names := range counts {
		if len(names) > 1 {
			fmt.Printf("%v\t%s\n", names, line)
		}
	}
}

func countLinesFromFile(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		//fmt.Printf("filename: %s, text: [%s]\n", f.Name(), input.Text())
		if slices.Contains(counts[input.Text()], f.Name()) == false {
			counts[input.Text()] = append(counts[input.Text()], f.Name())
		}
	}
}

/* test result
user@AL02279337 ch01 % go run ex1.4.go dup1 dup2
[dup1 dup2]		"fmt"
[dup1 dup2]		for line, n := range counts {
[dup1 dup2]
[dup1 dup2]	func main() {
[dup1 dup2]	import (
[dup1 dup2]	)
[dup1 dup2]		for input.Scan() {
[dup1 dup2]			}
[dup1 dup2]			counts[input.Text()]++
[dup1 dup2]	package main
[dup1 dup2]		"bufio"
[dup1 dup2]		"os"
[dup1 dup2]			if n > 1 {
[dup1 dup2]				fmt.Printf("%d\t%s\n", n, line)
[dup1 dup2]	}
[dup1 dup2]		counts := make(map[string]int)
[dup1 dup2]		}
*/
