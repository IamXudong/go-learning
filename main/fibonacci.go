// Fibonacci count the nth Fibonacci number.
package main

import (
	"flag"
	"fmt"
	"os"
)

var n = flag.Int("n", 1, "position of Fibonacci numbers")

func main() {
	flag.Parse()

	if *n < 1 {
		fmt.Println("n must be a integer number great or equal to 1")
		os.Exit(1)
	}

	fmt.Printf("The %dth Fibonacci number is: %d\n", *n, fib(*n))
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
