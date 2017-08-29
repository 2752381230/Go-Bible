package main

import (
    "fmt"
)

func main() {
    data := []string{"one", "", "three"}
    fmt.Printf("%q\n", data)
    fmt.Printf("%q\n", remove(data, 1))
    fmt.Printf("%q %d\n\n", data, len(data))

    data2 := []string{"one", "", "three"}
    fmt.Printf("%q\n", data2)
    data2 = remove(data2, 1)
    fmt.Printf("%q %d\n\n", data2, len(data2))

    data3 := []string{"one", "", "three"}
    fmt.Printf("%q\n", data3)
    copy(data3, remove(data3, 1))
    fmt.Printf("%q %d\n\n", data3, len(data3))

    data4 := []string{"one", "", "three"}
    fmt.Printf("%q\n", data4)
    data5 := remove(data4, 1)
    copy(data4, data5)
    fmt.Printf("%q %d\n\n", data4, len(data4))
    fmt.Printf("%q %d\n\n", data5, len(data5))
}

// in this studaton, we can see that `slice` 
// 1 copy to the function, out-function origin
//   slice also change
// 2 = will change the origin struct 
// 3 copy will not change the origin struct info
func remove(slice []string, i int) []string {
    copy(slice[1:], slice[1+1:])
    return slice[:len(slice)-1]
}
/*
0 1 2 3 4 5
a b c d e f

=> 1

0 1 2 3 4 5
a f c d e
*/
func remove2(slice []string, i int) []string {
    slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}
