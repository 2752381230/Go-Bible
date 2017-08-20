package main

import (
    "fmt"
    "os"
)

func main() {
    fname := "local_global_s5.go"
    if f, err := os.Open(fname); err != nil { // compile error: unused: f
        fmt.Println(err)
    }
    //f.ReadByte() // compile error: undefined f
    //f.Close()    // compile error: undefined f
    fmt.Println(f, "compile error: undefined f")
    /*
    f, err := os.Open(fname);
    if err != nil {
        fmt.Println(err)
    }
    f.ReadByte()
    f.Close()
    */
}
