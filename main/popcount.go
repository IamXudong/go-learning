package main

import (
	"fmt"
	"go-learning/popcount"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		count := popcount.PopCount(uint64(x))
		fmt.Printf("%d popcount: %d\n", x, count)
	}
}
