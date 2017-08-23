package main

import (
    "fmt"
)

/*
ax^2+bx+c=Ay+B
*/
func main() {
    var x, y, z int64
    var a, b    int64
    for x = 0; x < 256; x++ {
        for y = 0; y < 256; y++ {
            for z = 0; z < 256; z++ {
                tmp := x*256*256 + y*256 + z
                a = tmp%4096
                b = (tmp - a)/4096
                fmt.Println("(",x, y, z,") => (", a, b, ")")
            }
        }
    }
    for a = 0; a < 4096; a++ {
        for b = 0; b < 4096; b++ {
            tmp := a+b*4096
            z   = tmp%256
            tmp = (tmp-z)/256
            y   = tmp%256
            x   = (tmp-y)/256
            fmt.Println("(",a,b,") => (", x, y, z,")")
        }
    }
}
