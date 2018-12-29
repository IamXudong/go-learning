// 反转slice
package main

import "fmt"

var a = [...]int{0, 1, 2, 3, 4, 5}

func main() {
	fmt.Println(a)

	// 反转a [5, 4, 3, 2, 1, 0]
	revert(a[:])
	fmt.Println(a)

	// 重置a&a左移2位 [2, 3, 4, 5, 0, 1]
	revert(a[:])
	revert(a[:2])
	revert(a[2:])
	revert(a[:])
	fmt.Println(a)

	// a右移2位 [0, 1, 2, 3, 4, 5]
	revert(a[:])
	revert(a[:2])
	revert(a[2:])
	fmt.Println(a)
}

// 反转一个int类型的slice
func revert(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
