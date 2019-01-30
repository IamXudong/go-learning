package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/geometry2"
)

func main() {
	path := geometry2.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	p := geometry2.Point{2, 4}
	path.TranslateBy(p, true)
	fmt.Println(path)
}
