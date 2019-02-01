package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/counter"
)

func main() {
	var c counter.Counter
	fmt.Println(c.N())
	c.Increment()
	c.Increment()
	fmt.Println(c.N())
	c.Reset()
	fmt.Println(c.N())
}
