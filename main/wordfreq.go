// 统计文本文件中每个单词出现的次数
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var words = make(map[string]int)	// 单词计数
var puncs = make(map[string]int)	// 标点计数

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		word := ""
		for _, r := range line {
			if unicode.IsLetter(r) {
				word += string(r)
				continue
			}
			puncs[string(r)]++
			if word != "" {
				words[word]++
				word = ""
			}
		}
		if word != "" {
			words[word]++
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Words count: ")
	for w, c := range words {
		fmt.Printf("%6d | %-30s\n", c, w)
	}
	fmt.Println("\nPuncs count: ")
	for p, c := range puncs {
		fmt.Printf("%6d | %-30s\n", c, p)
	}
}
