package main

import (
	"fmt"
	"os"
)

func main() {
	_ = double(4)
	fmt.Println(triple(4))
	for _, filename := range os.Args[1:] {
		if err := doFile(filename); err != nil {
			fmt.Fprintf(os.Stderr, "doFile(%s): %v\n", filename, err)
		}
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Printf("%v\n", f)
	return nil
}
