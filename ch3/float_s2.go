package main

import (
    "fmt"
    "math"
)

func main() {
    var z float64
    // 下面的输出不需要 math
    fmt.Println(z, -z, 1/z, -1/z, z/z)
    // 下面的输出需要 math, ⚠️结果
    nan := math.NaN()
    fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false
}
