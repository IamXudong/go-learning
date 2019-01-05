// 翻转UTF-8编码字符串中的元素
// 传入参数为字符串对应字节的slice类型([]byte)
// 要求：就地修改(不重新分配内存)
package main

import (
	"fmt"
	"os"
)

var b []byte

func main() {
	if len(os.Args) > 1 {
		b = []byte(os.Args[1])
	}

	fmt.Printf("原字符串: %s\n", string(b))
	reverse(b)
	fmt.Printf("翻转结果: %s\n", string(b))
}

// 反转utf8 []byte slice字符
func reverse(b []byte) {
	if len(b) == 0 {
		return
	}

	for remain, size := len(b), runeSize(b[0]); remain > size; {
		// utf8序列解析出错
		if size == 0 {
			os.Exit(1)
		}

		// 当前字符置换至末尾
		rotate(b[:remain], size)

		// 设置下次置换参数
		remain -= size
		size = runeSize(b[0])
	}
}

// 就地旋转 []byte slice
func rotate(b []byte, step int) {
	l := len(b) // slice 长度

	// 长度或步长为0，直接返回
	if step == 0 || l == 0 {
		return
	}

	step %= l // 过滤无效旋转

	// 统一转换右旋转为左旋转
	if step < 0 {
		step += l
	}

	// 依次交换当前元素与步长偏移的元素
	i := 0
	for i+step < l {
		b[i], b[i+step] = b[i+step], b[i]
		i++
	}

	// 递归调用，直至步长前元素移至末尾
	if l%step != 0 {
		rotate(b[i:], step-l%step)
	}
}

// 计算utf8 []byte slice第一个字符的字节长度
func runeSize(b byte) int {
	z := byte(1 << 7)
	count := 0
	for b&z != 0 {
		count++
		z >>= 1
	}
	// 非字符起始字节
	if count == 1 {
		return 0
	}
	// ASII
	if count == 0 {
		return 1
	}

	return count
}
