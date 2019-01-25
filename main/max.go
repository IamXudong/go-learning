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
	maxium, err := max(vals...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "max: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Max", vals, ": ", maxium)
}

func max(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("need at least one integer")
	}
	maxium := vals[0]
	for _, val := range vals[1:] {
		if val > maxium {
			maxium = val
		}
	}
	return maxium, nil
}
