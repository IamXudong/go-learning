// 字符串替换函数
package main

import (
	"fmt"
	"os"
	"unicode"
)

var data = map[string]string{"name": "Steve", "country": "China", "job": "programer"}

func main() {
	s := os.Args[1]
	rst := expand(s, star)
	fmt.Println(rst)
}

func expand(s string, f func(s string) string) string {
	var flag = false
	var wi int
	var rst string
	for i, r := range s {
		if !flag && r == '$' {
			rst += string(s[wi:i])
			flag = true
			wi = i + 1
			continue
		}
		if flag && !unicode.IsLetter(r) {
			rst += f(string(s[wi:i]))
			flag = false
			wi = i
			continue
		}
	}
	if !flag && wi < len(s) {
		rst += string(s[wi:])
	}
	if flag && wi < len(s) {
		rst += f(string(s[wi:]))
	}
	return rst
}

func star(s string) string {
	rst, ok := data[s]
	if !ok {
		rst = "Unknow"
	}
	return rst
}
