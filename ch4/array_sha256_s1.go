package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	arr := [32]byte{'a', 'b', 'c'}
	fmt.Printf("%x\n", arr)
	zero(&arr)
	fmt.Printf("%x\n", arr)
}
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func zeroV2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
