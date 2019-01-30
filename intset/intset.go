// Package intset 提供一个非负整数集合
// 及相关操作方法
package intset

import (
	"bytes"
	"fmt"
)

// IntSet 是一个包含非负整数的集合
// 零值代表空的集合
type IntSet struct {
	Words []uint64
}

// Has 方法返回值表示是否存在非负整数 x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.Words) && s.Words[word]&(1<<bit) != 0
}

// Add 方法添加非负整数 x 到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.Words) {
		s.Words = append(s.Words, 0)
	}
	s.Words[word] |= 1 << bit
}

// UnionWith 将会对 s 和 t 做并集，并将结果存在 s 中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.Words {
		if i < len(s.Words) {
			s.Words[i] |= tword
		} else {
			s.Words = append(s.Words, tword)
		}
	}
}

// String 方法以字符串 "{1 2 3}" 的形式返回集中
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.Words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
