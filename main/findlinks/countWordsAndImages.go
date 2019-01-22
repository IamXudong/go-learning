// 获取页面，并统计页面上的单词和图片
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"unicode"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := countWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countWordsAndImages %s: %v\n", url, err)
			continue
		}
		fmt.Printf("Total: %d words, %d images: \n", len(words), len(images))
		for _, w := range words {
			fmt.Println(w)
		}
		fmt.Println("------------------------")
		for _, i := range images {
			fmt.Println(i)
		}
	}
}

func countWordsAndImages(url string) (words, images []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("gettting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = count(doc)
	return
}

func count(n *html.Node) (words, images []string) {
	if n.Data == "script" || n.Data == "style" {
		return
	}
	if n.Type == html.TextNode {
		down := false
		s := 0
		for i, r := range n.Data {
			if unicode.IsLetter(r) {
				if !down {
					s = i
				}
				down = true
				continue
			}
			if down {
				words = append(words, n.Data[s:i])
				down = false
			}
		}
		if down {
			words = append(words, n.Data[s:])
		}
	}
	if n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				images = append(images, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordsn, imagesn := count(c)
		words = append(words, wordsn...)
		images = append(images, imagesn...)
	}
	return
}
