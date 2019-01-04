// slice就地处理，合并UTF-8字节slice中所有相邻空白字符为一个ASCII空白字符
package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("")
	if len(os.Args) > 1 {
		b = []byte(os.Args[1])
	}
	fmt.Println(string(b))
	fmt.Println(string(combspace(b)))
}

func combspace(b []byte) []byte {
	s := string(b)
	n := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		// 空字符合并
		if unicode.IsSpace(r) {
			if n < 1 || b[n-1] != ' ' {
				b[n] = ' '
				n++
			}
			continue
		}
		// 非空字符直接copy
		copy(b[n:n+size], b[i-size:i])
		n += size
	}
	return b[:n]
}
