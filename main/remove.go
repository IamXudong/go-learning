// 移除整数slice中第i个元素
package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("s\t%v\n", s)
	r1 := remove1(s, 2)
	fmt.Printf("remove1(s, 2)\t%v\n", r1)
	r2 := remove1(r1, 2)
	fmt.Printf("remove2(r1, 2)\t%v\n", r2)
}

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	// 弹出栈里最前面的一个元素
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
