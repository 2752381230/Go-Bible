package main

import (
    "fmt"
)

func f() {}

var g = "g"

func main() {
    f := "f"
    fmt.Println("This f is locale", f)

    fmt.Println("This g is global", g)

    fmt.Println("This h is undefined")
    //fmt.Println("This h is undefined", h)
}
