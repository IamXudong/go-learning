// comma向表示十进制非负整数的字符串中插入逗号
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	fmt.Printf("%s\t=\t%s\n", "commaA(\""+s+"\")", commaA(s))
	fmt.Printf("%s\t=\t%s\n", "commaB(\""+s+"\")", commaB(s))
	fmt.Printf("%s\t=\t%s\n", "commaC(\""+s+"\")", commaC(s))
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
		buf.WriteString(s[n : n+3])
		n += 3
	}
	return buf.String()
}

// 支持小数和符号的comma函数
func commaC(s string) string {
	var buf bytes.Buffer
	begin, end := 0, len(s) // 整数部分的开始与结束位置

	// 更新整数部分开始位置
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		begin = 1
	}

	// 更新整数部分结束位置
	dot := strings.Index(s, ".")
	if dot > 0 {
		end = dot
	}

	buf.WriteString(s[:begin]) // 写入符号部分
	intSec := s[begin:end]     // 对整数部分创建索引

	// 写入整数部分
	n := len(intSec) % 3
	buf.WriteString(intSec[:n])
	for n <= len(intSec)-3 {
		if n > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(intSec[n : n+3])
		n += 3
	}

	buf.WriteString(s[end:]) // 写入小数部分
	return buf.String()
}
