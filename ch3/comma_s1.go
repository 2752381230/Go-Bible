package main

import (
    "fmt"
    "strings"
)

func main() {
    strs := []string {"12", "123", "1234", "12345", "123456", "1234567", "12345678"}

    for _, str := range strs {
        fmt.Println(str, " => ", comma(str))
    }
}
// 很直观的方式 head -> tail
func comma1(s string) string {
    if len(s) <= 3 {
        return s
    }
    lens := len(s)
    left := lens%3
    str  := make([]string, 0)
    if left != 0 {
        str = append(str, s[0:left])
    }
    for i := left; i < lens; i = i + 3 {
        str = append(str, s[i:i+3])
    }
    return strings.Join(str, ",")
}
// 逆序的方式,  tail-> head
func comma2(s string) string {
    if len(s) <= 3 {
        return s
    }
    lens  := len(s)
    index := lens
    left  := lens%3
    for i := lens-1; i >= 0; i-- {
        if i %3 == left {
            fmt.Println(s[i:index])
            index = i
        } else if i == 0 {
            fmt.Println(s[i:index])
        }
    }
    return ""
}
// 递归的方式
func comma(s string) string {
    n := len(s)
    if n <= 3 {
        return s
    }
    return comma(s[:n-3]) + "," + s[n-3:]
}
