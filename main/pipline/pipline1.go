package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; ; x++ {
			if x > 10 {
				break
			}
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer (在主 goroutine 中)
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
