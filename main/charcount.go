// charcount 计算Unicode字符的个数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode 字符数量
	var utflen [utf8.UTFMax + 1]int // UTF-8 编码的长度
	invalid := 0                    // 非法 UTF-8 字符数量
	letter := 0                     // 字母数量
	number := 0                     // 数字数量
	space := 0                      // 空字符数量

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 返回 rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsNumber(r) {
			number++
		}
		if unicode.IsSpace(r) {
			space++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("\nlen\tcount")
	for i, n := range utflen {
		if n > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\nletters\t%d\n", letter)
	fmt.Printf("numbers\t%d\n", number)
	fmt.Printf("spaces\t%d\n", space)
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
