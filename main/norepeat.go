// 就地处理，去除[]string slice中相邻重复字符串元素
package main

import "fmt"

func main() {
	s1 := []string{}
	s2 := []string{"one"}
	s3 := []string{"one", "two", "two", "three", "four", "five", "five", "six"}
	fmt.Println(s1, s2, s3)
	s1 = norepeat(s1)
	s2 = norepeat(s2)
	s3 = norepeat(s3)
	fmt.Println(s1, s2, s3)
}

func norepeat(s []string) []string {
	n := 0
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] != s[n-1] {
			s[n] = s[i]
			n++
		}
	}
	return s[:n]
}
