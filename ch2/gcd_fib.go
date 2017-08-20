package main

import "fmt"

func main() {
	x, y := 15, 5
	fmt.Println(x," ",y," 最大公约数是",gcd(x,y))
	fmt.Printf("第%d项的斐波那契额数列是:", x)
	fmt.Println(fib(x))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
