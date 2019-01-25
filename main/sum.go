package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var vals []int
	for _, arg := range os.Args[1:] {
		val, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "strconv: %v\n", err)
			os.Exit(1)
		}
		vals = append(vals, val)
	}
	total := sum(vals...)
	fmt.Println("sum", vals, ": ", total)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
