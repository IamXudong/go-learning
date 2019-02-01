// Package intset提供一个非负整数集合
// 及相关操作方法
package intset

import (
	"bytes"
	"fmt"
)

const bits = 32 << (^uint(0) >> 63)

// IntSet是一个包含非负整数的集合
// 零值代表空的集合
type IntSet struct {
	words []uint
}

// Has方法返回值表示是否存在非负整数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/bits, uint(x%bits)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add方法添加非负整数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/bits, uint(x%bits)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith将会对s和t做并集，并将结果存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String方法以字符串"{1 2 3}"的形式返回集中
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len方法返回s集合中元素的个数
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		temp := word
		for temp != 0 {
			count++
			temp &= temp - 1
		}
	}
	return count
}

// Remove方法从集合中去除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/bits, uint(x%bits)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// Clear方法删除所有元素
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy方法返回s的副本c
func (s *IntSet) Copy() *IntSet {
	c := new(IntSet)
	c.words = make([]uint, s.Len())
	copy(c.words, s.words)
	return c
}

// AddAll方法添加多个非负整数ints到集合中
func (s *IntSet) AddAll(ints ...int) {
	for _, x := range ints {
		s.Add(x)
	}
}

// Intersec方法返回s和t的交集
func (s *IntSet) Intersec(t *IntSet) *IntSet {
	inter := new(IntSet)
	for i, sword := range s.words {
		if i >= len(t.words) {
			break
		}
		inter.words = append(inter.words, sword&t.words[i])
	}
	return inter
}

// Diff方法返回s和t的差集
func (s *IntSet) Diff(t *IntSet) *IntSet {
	diff := new(IntSet)
	for i, sword := range s.words {
		if i < len(t.words) {
			diff.words = append(diff.words, sword&^t.words[i])
		} else {
			diff.words = append(diff.words, sword)
		}
	}
	return diff
}

// SymDiff方法返回s和t的对称差
func (s *IntSet) SymDiff(t *IntSet) *IntSet {
	symdiff := new(IntSet)
	lwlen, swlen := len(s.words), len(t.words)
	lset, sset := s, t
	if swlen > lwlen {
		lset, sset = t, s
		lwlen, swlen = swlen, lwlen
	}
	for i, word := range lset.words {
		if i < swlen {
			symdiff.words = append(symdiff.words, word^sset.words[i])
		} else {
			symdiff.words = append(symdiff.words, word)
		}
	}
	return symdiff
}

// Elems方法返回集合的元素slice
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, bits*i+j)
			}
		}
	}
	return elems
}
