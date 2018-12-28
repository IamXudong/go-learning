package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("YB/ZB: %d\n", YB/ZB)
}
