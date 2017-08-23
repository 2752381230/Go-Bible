package main

import "fmt"

func main() {
    //d o x 控制输出的进制格式
    //[1] 表示使用第一个操作数，以此类推
    //# 表示在输出时生成前导 0 0x 0X 
    o := 0666
    fmt.Printf("%d %[1]o %#[1]o\n", o)
    x := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
