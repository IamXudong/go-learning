// Cfk converts its nemberic arguments to Celsius, Fahrenheit and Kelvin.
package main

import (
	"fmt"
	"go-learning/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		k := tempconv.Kelvin(t)

		fmt.Printf(
			"%s = %s\t%s = %s\n%s = %s\t%s = %s\n%s = %s\t%s = %s\n",
			c, tempconv.CToF(c), f, tempconv.FToC(f),
			k, tempconv.KToC(k), c, tempconv.CToK(c),
			f, tempconv.FToK(f), k, tempconv.KToF(k))
	}
}
