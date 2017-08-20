package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] { // Args will parse space on the command line
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
