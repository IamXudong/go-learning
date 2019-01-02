// 一次遍就地历旋转slice元素
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var n int
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println(s)
	rotate1(s, n)
	fmt.Println(s)
	rotate2(s, -n)
	fmt.Println(s)
}

// 不使用递归
func rotate1(s []int, n int) {
	var step int
	step = n % len(s) // 转化无效旋转

	// 转化右旋转为左旋转
	if step < 0 {
		step = len(s) + step
	}

	stop := 0
	for i := 0; i < len(s)-1; i++ {
		if i+step > len(s)-1 {
			r := (i - stop) % step
			if r == 0 {
				break
			}
			step = step - r
			stop = i
		}
		j := i + step
		s[i], s[j] = s[j], s[i]
	}
}

// 使用递归
func rotate2(s []int, step int) {
	step %= len(s)
	if step == 0 {
		return
	}

	// 转化右旋转为左旋转
	if step < 0 {
		step = len(s) + step
	}

	i := 0
	for i+step < len(s) {
		s[i], s[i+step] = s[i+step], s[i]
		i++
	}
	step = step - i%step
	rotate2(s[i:], step)
}
