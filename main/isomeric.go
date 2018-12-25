// isomeric函数检测两个字符串是否为同文异构
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Printf("%s\t=\t%v\n", "isomeric("+s1+", "+s2+")", isomeric(s1, s2))
}

// 遍历检索s1中的字节是否存在于s2中
func isomeric(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if strings.Index(s2, s1[i:i+1]) < 0 {
			return false
		}
	}
	return true
}
