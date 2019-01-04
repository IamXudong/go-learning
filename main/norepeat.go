// 就地处理，去除[]string slice中相邻重复字符串元素
package main

import "fmt"

func main() {
	s := []string{"one", "two", "two", "three", "four", "five", "five", "six"}
	fmt.Println(s)
	s = norepeat(s)
	fmt.Println(s)
}

func norepeat(s []string) []string {
	if len(s) < 1 {
		return s
	}
	n := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[n] {
			n++
			s[n] = s[i]
		}
	}
	return s[:n+1]
}
