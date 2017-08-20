package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

// add read info from files
func main() {
	counts := make(map[string]int)
	// e1.4
	record := make(map[string]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			record[line] = filename
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s in files: %s\n", n, line, record[line])
	}
}
