package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown, Press return to abourt.")
	select {
	case <-time.After(10 * time.Second):
		// 不执行任何操作
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
