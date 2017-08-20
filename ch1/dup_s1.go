package main

import (
	"fmt"
	"os"
	"bufio" // buffer io
)

func main() {
	fmt.Println("crel+d to stop the input")
	counts := make(map[string]int)
	input  := bufio.NewScanner(os.Stdin)
	for input.Scan() { // read line and remove \r\n at the end of the line
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
