package main

import (
    "fmt"
    "strings"
    "path/filepath"
)

func main() {
    strs := []string {"/ab/b/cd.go", "a.b.c.d.go", "abc"}

    for _, str := range strs {
        fmt.Println(str, " => ", basename(str))
    }
}
func basename(s string) string {
    sanddot := filepath.Base(s)
    if dot := strings.LastIndex(sanddot, "."); dot >= 0 {
        s = sanddot[:dot]
    }
    return s
}
