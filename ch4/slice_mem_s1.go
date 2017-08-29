package main

import (
    "fmt"
)

func main() {
    data := []string{"one", "", "three"}
    fmt.Printf("%q\n", data)
    fmt.Printf("%q\n", nonempty(data))
    fmt.Printf("%q\n\n", data)
    data2 := []string{"one", "", "three"}
    fmt.Printf("%q\n", data2)
    fmt.Printf("%q\n", nonempty2(data2))
    fmt.Printf("%q\n\n", data2)
}

// 两个函数实现的效果一致
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}
func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}
