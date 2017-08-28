package main

import (
    "fmt"
    "strings"
)

func main() {
    strs := []string {"/ab/b/cd.go", "a.b.c.d.go", "abc"}

    for _, str := range strs {
        fmt.Println(str, " => ", basename(str))
    }
}
func basename(s string) string {
    slash := strings.LastIndex(s, "/")
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}
