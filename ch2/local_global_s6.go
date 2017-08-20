package main

import (
    "fmt"
    "os"
)
var cwd string

func init() {
    // := 导致 cwd 在内部被视为 loca
    cwd, err := os.Getwd() // compile error: unused: cwd
    if err != nil {
        fmt.Printf("os.Getwd failed: %v", err)
    }
    // 这个地方的打印更隐藏了 cwd 是init 内 var 的事实
    fmt.Printf("working dir = %v", cwd)
}

func main() {
    fmt.Println("main ", cwd)
}
