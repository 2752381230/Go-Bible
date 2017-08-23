package main

import (
    "fmt"
    "math"
)

func main() {
    //g 打印紧凑的浮点数 
    //e f 平常使用
    for x := 0; x < 8; x++ {
        fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
    }
}
