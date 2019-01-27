package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/geometry"
)

func main() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	fmt.Println(geometry.PathDistance(perim))
	r := geometry.Point{1, 2}
	rptr := &r
	(&r).Scale(2)
	fmt.Println(r)
	rptr.Scale(2)
	fmt.Println(r)
	fmt.Println(rptr.Distance(geometry.Point{4, 8}))
	fmt.Println(*rptr)

}
