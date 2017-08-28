package main

import (
    "fmt"
)

func main() {
    strs := []string {"/ab/b/cd.go", "a.b.c.d.go", "abc"}

    for _, str := range strs {
        fmt.Println(str, " => ", basename(str))
    }
}
func basename(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }
    return s
}
