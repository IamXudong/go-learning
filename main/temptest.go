package main

import (
	"fmt"
	"go-learning/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))

	s := tempconv.FToC(212.0)
	fmt.Println(s.String())
	fmt.Printf("%v\n", s)
	fmt.Printf("%s\n", s)
	fmt.Println(s)
	fmt.Printf("%g\n", s)
	fmt.Println(float64(s))
}
