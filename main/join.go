package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "join: need at least 1 param for separator\n")
		os.Exit(1)
	}
	sep := os.Args[1]
	vals := os.Args[2:]
	fmt.Printf("Join %v with %s: %s\n", vals, sep, join(sep, vals...))
}

func join(sep string, vals ...string) string {
	var rst, sepn string
	for _, val := range vals {
		rst += sepn + val
		sepn = sep
	}
	return rst
}
