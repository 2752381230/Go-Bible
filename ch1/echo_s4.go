package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// e1.1
	fmt.Println(strings.Join(os.Args[0:], " "))
	// e1.2
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
	// e1.3
}
