package main

import (
	"flag"
	"fmt"
	"github.com/stevzhang01/go-learning/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temprature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
