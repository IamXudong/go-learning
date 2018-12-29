// 比较两个字符串slice是否相等
package main

import "fmt"

var a = []string{"a", "b", "c"}
var b = []string{"a", "c", "e"}
var c = []string{"a", "b", "c"}

func main() {
	fmt.Printf("%s\t%v\n", "a", a)
	fmt.Printf("%s\t%v\n", "b", b)
	fmt.Printf("%s\t%v\n", "c", c)

	fmt.Printf("%s\t%t\n", "equal(a, b)", equal(a, b))
	fmt.Printf("%s\t%t\n", "equal(a, c)", equal(a, c))
}

// 检测两个字符串slice是否相等
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
