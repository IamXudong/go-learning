// Package writecounter 实现io.Writer接口方法
// 提供字符串数量统计相关方法
package writecounter

import (
	"bufio"
	"bytes"
	"io"
)

// ByteCounter统计字节数量
type ByteCounter int

// ByteCounter实现io.Writer接口
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // 转换 Int 为 ByteCounter 类型
	return len(p), nil
}

// LineCounter统计行数量
type LineCounter int

// LineCounter实现io.Writer接口
func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*l++
	}
	return len(p), nil
}

// WordsCounter统计单词数量
type WordsCounter int

// WordsCounter实现io.Writer接口
func (w *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return len(p), nil
}

// CountWriter封装一个io.Writer
// 并统计写入的字节数量
type CountWriter struct {
	w     io.Writer
	count int64
}

// CountWriter实现io.Writer接口
func (cw *CountWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.count += int64(n)
	return n, err
}

// CountingWriter输入一个io.Writer
// 输出一个封装了输入值的新Writer
// 以及一个指向int64的指针
// 该指针对应的值是新的Writer写入的字节数
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw CountWriter
	cw.w = w
	return &cw, &cw.count
}
