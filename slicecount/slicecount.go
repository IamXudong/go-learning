// Package slicecount 使用map来记录一个字符串slice被add函数调用的次数
package slicecount

import "fmt"

var count = make(map[string]int)

// key 生成slice对应到map的string类型key
func key(list []string) string { return fmt.Sprintf("%q", list) }

// Add 添加次数
func Add(list []string) { count[key(list)]++ }

// Count 返回次数
func Count(list []string) int { return count[key(list)] }
