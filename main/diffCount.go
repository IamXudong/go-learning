// 统计两个sha256散列不同的位数
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

// 8位整数置位个数速查表
func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	c1, c2, count := diffCount(s1, s2)
	fmt.Printf("%s\t%x\n%s\t%x\nDifferent count: %d\n", s1, c1, s2, c2, count)
}

// 计算两个字符串sha256散列的不同位数
func diffCount(s1, s2 string) ([32]uint8, [32]uint8, byte) {
	var count byte
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	for i, _ := range c1 {
		v := c1[i] ^ c2[i] // 不同位的值
		n := pc[v]         // 不同位的个数(快查表方法)
		//      n := popCount(v)   // 不同位的个数(低位清零法)
		count += n
	}
	return c1, c2, count
}

// 低位清零法计算一个整数被置位的个数
func popCount(v uint8) byte {
	var count byte
	for v > 0 {
		v = v & (v - 1)
		count++
	}
	return count
}
