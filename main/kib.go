package main

import (
	"fmt"
)

const (
	_ = 1 << (10 * iota)
	KIB
	MIB
	GIB
	TIB
	PIB
	EIB
	ZIB
	YIB
)

func main() {
	fmt.Printf("KIG: %d\n", KIB)
	fmt.Printf("GIB: %d\n", GIB)
	fmt.Printf("YIB/ZIB: %d\n", YIB/ZIB)
}
