package main

import "fmt"

func f() int {
    return 1
}
func g(x int) int {
    return x
}

func main() {
    if x := f(); x == 0 {
            fmt.Println(x)
    } else if y := g(x); x == y {
            fmt.Println(x, y)
    } else {
            fmt.Println(x, y)
    }
    //fmt.Println(x, y) // compile error: x and y are not visible here
    //第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问
}
