// 解析文件路径中的文件名
package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	s := os.Args[1]
	fmt.Printf("%s\t=\t%s\n", "basenameA(s)", basenameA(s))
	fmt.Printf("%s\t=\t%s\n", "basenameB(s)", basenameB(s))
}

// 不使用strings包方式
func basenameA(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

// 使用strings包方式
func basenameB(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}

	return s
}
