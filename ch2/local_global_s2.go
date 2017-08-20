package main

import "fmt"

func main() {
    x := "hello!"                // main中的局部var
    for i := 0; i<len(x); i++ {
        x := x[i]               // x for内部的var, x[i]main的var
        if x!= '!' {
            x := x + 'A' - 'a'  // x for内部的var
            fmt.Printf("%c", x) // HELLO
        }
    }
}
