// Gcd compute the greatest common divisor of two integer x and y.
package main

import (
	"flag"
	"fmt"
)

var x = flag.Int("x", 1, "number x")
var y = flag.Int("y", 0, "number y")

func main() {
	flag.Parse()
	fmt.Printf("the gcd of number %d and %d is: %d\n", *x, *y, gcd(*x, *y))
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
