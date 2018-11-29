package main

import (
	"fmt"
	"go-learning/tempconv"
)

func main() {
	fmt.Printf("%v\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%v\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	boilingK := tempconv.CToK(tempconv.BoilingC)
	fmt.Printf("%v\n", boilingK-tempconv.CToK(tempconv.FreezingC))
	// fmt.Printf("%g\n", boilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	var k tempconv.Kelvin
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))

	c = tempconv.FToC(212.0)
	k = tempconv.CToK(-273.15)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	fmt.Println(k.String())
}
