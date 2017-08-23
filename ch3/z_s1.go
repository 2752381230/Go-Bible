package main

import (
	"fmt"
)

func main() {
	var x complex128 = complex(1,2)
	var y complex128 = complex(3,4)
	var z = 1+2i

	fmt.Println("x*y = ", x*y)
	fmt.Println("real(x*y) = ", real(x*y))
	fmt.Println("imag(x*y) = ", imag(x*y))
	fmt.Println(1i*1i)
	fmt.Println("x*z = ", x*z)
	fmt.Println("real(x*z) = ", real(x*z))
	fmt.Println("imag(x*z) = ", imag(x*z))
}
