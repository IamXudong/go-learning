// 判断两个map拥有相同的键和值
package main

import "fmt"

func main() {
	x := map[string]int{"A": 0, "B": 1, "C": 2}
	y := map[string]int{"A": 1, "B": 2, "C": 5}
	z := map[string]int{"A": 0, "B": 1, "C": 2}

	fmt.Println("equal(x, y):\t", equal(x, y))
	fmt.Println("equal(x, z):\t", equal(x, z))
}

func equal(x map[string]int, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || xv != yv {
			return false
		}
	}
	return true
}
