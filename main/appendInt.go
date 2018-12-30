// appendInt函数为int slice追加元素
package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

// 追加元素到int slice
func appendInt(x []int, i int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice仍有增长空间，扩展slice内容
		z = x[:zlen]
	} else {
		zcap := zlen
		// slice已无空间，为它新分配一个底层数组
		// 为了分摊线性复杂性，容量扩展一倍
		if zcap <= 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = i
	return z
}
