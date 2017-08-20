package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	/*
	var xx type  must be have type, recommend this format
	s := "" it's ok, 
	var s = ""  No
	var s string = ""  No
	*/
	for i := 1; i < len(os.Args); i++ { // Args will parse space on the command line
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
