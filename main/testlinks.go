package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/links"
	"os"
)

func main() {
	url := os.Args[1]
	alinks, err := links.Extract(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "links.Extract: %v\n", err)
		os.Exit(1)
	}
	for _, link := range alinks {
		fmt.Println(link)
	}
}
