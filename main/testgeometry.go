package main

import (
	"fmt"
	"github.com/stevzhang01/go-learning/geometry"
	"image/color"
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
	red := color.RGBA{0xff, 0x00, 0x00, 0x00}
	blue := color.RGBA{0x00, 0x00, 0xff, 0x00}
	p := geometry.CorloredPoint{
		geometry.Point{1, 1},
		red,
	}
	q := geometry.CorloredPoint{
		geometry.Point{5, 4},
		blue,
	}
	fmt.Println(p.Distance(q.Point))
	p.Scale(2)
	q.Scale(2)
	fmt.Println(p.Distance(q.Point))
	fmt.Println("------")
	p1 := geometry.Point{1, 2}
	q1 := geometry.Point{4, 6}
	distanceFromP := p1.Distance
	fmt.Println(distanceFromP(q1))
	var origin geometry.Point
	fmt.Println(distanceFromP(origin))
	scaleP := (&p1).Scale
	scaleP(2)
	fmt.Println(p1)
	scaleP(3)
	fmt.Println(p1)
	scaleP(10)
	fmt.Println(p1)
	fmt.Println("--------")
	p2 := geometry.Point{1, 2}
	q2 := geometry.Point{4, 6}
	distance := geometry.Point.Distance
	fmt.Println(distance(p2, q2))
	fmt.Printf("%T\n", distance)
	scale := (*geometry.Point).Scale
	scale(&p2, 2)
	fmt.Println(p2)
	fmt.Printf("%T\n", scale)
}
