// 键为slice的map实现方案
// 使用map来记录一个字符串slice被add函数调用的次数
package main

import (
	"fmt"
	"go-learning/slicecount"
)

func main() {
	x := []string{"a", "b", "c"}
	y := []string{"e", "f", "g"}
	z := []string{"h"}

	for i := 0; i < 10; i++ {
		slicecount.Add(x)
	}

	for i := 0; i < 2; i++ {
		slicecount.Add(y)
	}

	fmt.Println("x: ", slicecount.Count(x))
	fmt.Println("y: ", slicecount.Count(y))
	fmt.Println("z: ", slicecount.Count(z))
}
