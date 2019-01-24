package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/links"
	"log"
	"os"
)

func main() {
	// 开始广度遍历
	// 从命令行参数开始
	breathFirst(crawl, os.Args[1:])
}

// breathFirst 对每个 worklist 元素调用 f
// 并将返回的内容添加到 worklist 中，对每个元素，最多调用一次 f
// f is called at most once for each item.
func breathFirst(f func(url string) []string, worklist []string) {
	seen := make(map[string]bool)
	for worklist != nil {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
