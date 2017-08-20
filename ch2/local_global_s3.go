package main

import "fmt"

func main() {
    x := "hello"           // main中的局部var
    for _, x := range x {   // x for内部的， range x main的
        x := x + 'A' - 'a'  // x for内部的var
        fmt.Printf("%c", x) // HELLO
    }
}
