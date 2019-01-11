// graphic2d使用用graphic2d包
package main

import (
	"fmt"

	"github.com/stevzhang01/go-learning/graphic2d"
)

func main() {
	pp := &graphic2d.Point{3, 6}
	pp = graphic2d.PointScale(pp, 10)
	fmt.Println(pp.X, pp.Y)
	p := graphic2d.Point{1, 2}
	q := graphic2d.Point{X: 2, Y: 1}
	r := graphic2d.Point{X: 1, Y: 2}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)
	fmt.Println(p.X == r.X && p.Y == r.Y)
	fmt.Println(p == r)
	var s graphic2d.Wheel1
	fmt.Printf("%#v\n", s)
	fmt.Println(s.X, s.Y)
	s.X = 10
	fmt.Println(s)
	t := graphic2d.Wheel1{
		graphic2d.Circle1{
			graphic2d.Point{1, 2},
			5,
		},
		10,
	}
	fmt.Printf("%#v\n", t)
	fmt.Println(t.Circle1.Radius)
	fmt.Println(t.Circle1)
	fmt.Println(t.Radius)
	t.Radius = 100
	fmt.Println(t.Radius)
	var u graphic2d.Circle2
	u.X = 10
	fmt.Printf("%#v\n", u)
}
