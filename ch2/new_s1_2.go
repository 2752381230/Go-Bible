package main

import(
	"fmt"
)

func main() {
	n1, n2 := newInt1(), newInt2()
	*n1 = 1
	*n2 = 2
	fmt.Println(n1, n2, *n1, *n2)
	fmt.Println(newInt1() == newInt1())
}
func newInt1() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}
