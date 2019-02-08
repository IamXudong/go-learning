package tempconv

import (
	"flag"
	"fmt"
)

// *celsiusFlag satisfies the flag.Value interface
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var value float64
	var unit string
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFalg define a Celsius flag with the specified name,
// defalut name, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
