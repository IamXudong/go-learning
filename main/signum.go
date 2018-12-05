// Signum 返回一个整数的符号
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		x, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d\t%d\n", x, SigNum(x))
	}
}

func SigNum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
	return 0
}
