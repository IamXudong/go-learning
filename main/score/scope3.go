package main

import "fmt"

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
}

func f() int {
	return 1
}

func g(x int) int {
	x++
	return x
}
