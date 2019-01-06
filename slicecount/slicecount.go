// 使用map来记录一个字符串slice被add函数调用的次数
package slicecount

import "fmt"

var count = make(map[string]int)

func key(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string)        { count[key(list)]++ }
func Count(list []string) int  { return count[key(list)] }
