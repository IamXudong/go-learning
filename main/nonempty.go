// 移除字符串slice中的空元素
package main

import "fmt"

func main() {
	data1 := []string{"one", "", "three"}
	data2 := []string{"one", "", "", "four"}
	fmt.Printf("nonempty1(data1)\t%q\n", nonempty1(data1))
	fmt.Printf("data1\t%q\n", data1)
	fmt.Printf("nonempty2(data2)\t%q\n", nonempty2(data2))
	fmt.Printf("data2\t%q\n", data2)
}

// nonempty返回一个新的slice, slice中的元素都是非空字符串
// 在函数调用过程中， 底层数组的元素发生了改变
func nonempty1(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 使用append函数实现
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
