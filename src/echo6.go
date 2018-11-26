package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for index, value := range os.Args[1:] {
		s := strconv.Itoa(index) + " => " + value
		fmt.Println(s)
	}
}
