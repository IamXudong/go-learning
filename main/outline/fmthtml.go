// 解析并格式化输出HTML文档
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

var depth int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdout, "parsing: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, " ", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
			depth++
		}
	case html.TextNode:
		s := strings.TrimSpace(n.Data)
		if s != "" {
			fmt.Printf("%*s%s\n", depth*2, " ", s)
		}
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, " ", n.Data)
	default:
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
