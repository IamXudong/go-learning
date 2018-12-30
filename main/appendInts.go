// appendInts函数可一次为int slice追加多个元素
package main

import "fmt"

func main() {
	var x []int
	x = appendInts(x, 1)
	x = appendInts(x, 2, 3)
	x = appendInts(x, 4, 5, 6)
	x = appendInts(x, x...)
	fmt.Println(x)
}

// 为ints slice追加多个元素
func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < len(x)*2 {
			zcap = len(x) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
