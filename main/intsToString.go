// intsToString函数格式化int数组为字符串
package main

import (
	"fmt"
	"bytes"
)

func main() {
	fmt.Printf("%s\t=\t%s\n", "intsToString([]int{1, 2, 5})", intsToString([]int{1, 2, 5}))
}

// 格式化int数组为字符串
func intsToString(ints []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
