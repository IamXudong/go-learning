// 二叉树插入排序
package main

import (
	"fmt"
	"go-learning/treesort"
)

func main() {
	s := []int{5, 7, 9, 4, 1, 3, 2, 8, 6}
	fmt.Println("slice:\t", s)
	treesort.Sort(s)
	fmt.Println("排序:\t", s)
}
