// comma向表示十进制非负整数的字符串中插入逗号
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	fmt.Printf("%s\t=\t%s\n", "commaA(\""+s+"\")", commaA(s))
	fmt.Printf("%s\t=\t%s\n", "commaB(\""+s+"\")", commaB(s))
}

// 递归向字符串插入3位分隔逗号
func commaA(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaA(s[:n-3]) + "," + s[n-3:]
}

// 非递归方式向字符串中插入3位分隔逗号
func commaB(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	buf.WriteString(s[:n])

	for n <= len(s)-3 {
		if n > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[n:n+3])
		n += 3
	}
	return buf.String()
}
