package main

import "fmt"

func main() {
    //c 字符 
    //q 打印带有 '' 的字符
    ascii := 'a'
    unicode := '国'
    newline := '\n'
    fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
    fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
    fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'
}
