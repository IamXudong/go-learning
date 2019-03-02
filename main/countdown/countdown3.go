package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 创建中止通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown, Press return to abort.")

	// 创建一个 ticker 用于间隔计时
	// 并在主函数返回前关闭该 ticker
	// 避免没有 goroutine 接收 ticker 造成泄漏
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
			// 不执行任何操作
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
