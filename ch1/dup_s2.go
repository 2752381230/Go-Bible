package main

import (
	"fmt"
	"os"
	"bufio" // buffer io
)

// add read info from files
func main() {
	counts := make(map[string]int)
	files  := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("crel+d to stop the input")
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
