package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var depth int

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "outline2: %v\n", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Fprintf(os.Stdout, "gettting %s: %s\n", url, resp.Status)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stdout, "parsing %s as HTML: %V\n", url, err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
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
