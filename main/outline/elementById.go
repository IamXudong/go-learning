// 解析HTML文档并查找指定 ID 的 NODE
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var gid string

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stdout, "parsing: %v\n", err)
		os.Exit(1)
	}
	n := elementById(doc, "test")
	fmt.Println(n.Data)
	for _, a := range n.Attr {
		fmt.Println(a.Key, "=", a.Val)
	}
}

func elementById(n *html.Node, id string) *html.Node {
	gid = id
	return forEachNode(n, checkId, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r := forEachNode(c, pre, post)
		if r != nil {
			return r
		}
	}
	if post != nil {
		if post(n) {
			return n
		}
	}
	return nil
}

func checkId(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == gid {
				return true
			}
		}
	}
	return false
}
