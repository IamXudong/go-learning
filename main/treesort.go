// 二叉树插入排序
package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/treesort"
)

func main() {
	s := []int{5, 7, 9, 4, 1, 3, 2, 8, 6}
	fmt.Println("slice:\t", s)
	fmt.Print("序列值: \t")
	treesort.PrintValues(s)
	treesort.Sort(s)
	fmt.Println("排序:\t", s)
}
