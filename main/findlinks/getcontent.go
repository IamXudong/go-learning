package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Getcontent: %v\n", err)
		os.Exit(1)
	}
	getcontent(doc)
}

func getcontent(n *html.Node) {
	if n == nil {
		return
	}
	if n.Data == "script" || n.Data == "style" {
		getcontent(n.NextSibling)
		return
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	getcontent(n.FirstChild)
	getcontent(n.NextSibling)
}
