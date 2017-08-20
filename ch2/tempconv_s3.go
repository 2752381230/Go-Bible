package main

import (
    "fmt"
    "Go-Bible/ch2/tempconv"
    // 上面引用的时候注意GOPATH设置好,这个地方的包等同于dirname
    // mkdir Go-Bible/ch2/tempconv 
    // cp tempconv_s2.go conv_s1.go tempconv
    // 测试完成之后将 tempconv 删除掉了
    // tempconv_s3 不能放到其中，导致两个包在一个dir中的错误
)

func main() {
    fmt.Println("Kelvin is: %f", tempconv.AbsoluteZeroC)
    fmt.Println("To tran Kelvin to F")
    fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
}
