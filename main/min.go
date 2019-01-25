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
	minium, err := min(vals...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "min: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("min %v : %d\n", vals, minium)
}

func min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("need at least one integer")
	}
	minium := vals[0]
	for _, val := range vals[1:] {
		if val < minium {
			minium = val
		}
	}
	return minium, nil
}
