// Test package popcount.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/stevzhang01/go-learning/popcount"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	startA := time.Now()
	countA := popcount.PopCount(uint64(x))
	fmt.Printf("A\t%d\t%.12fs\n", countA, time.Since(startA).Seconds())

	startB := time.Now()
	countB := popcount.PopCountB(uint64(x))
	fmt.Printf("B\t%d\t%.12fs\n", countB, time.Since(startB).Seconds())

	startC := time.Now()
	countC := popcount.PopCountC(uint64(x))
	fmt.Printf("C\t%d\t%.12fs\n", countC, time.Since(startC).Seconds())

	startD := time.Now()
	countD := popcount.PopCountD(uint64(x))
	fmt.Printf("D\t%d\t%.12fs\n", countD, time.Since(startD).Seconds())
}
