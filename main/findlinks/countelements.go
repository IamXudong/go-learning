// 输出从标准输入中读入的 HTML 文档中的所有链接
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var elements = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Countelements: %v\n", err)
		os.Exit(1)
	}
	count(doc)
	for e, c := range elements {
		fmt.Printf("%8s\t%8d\n", e, c)
	}
}

func count(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	count(n.FirstChild)
	count(n.NextSibling)
}
