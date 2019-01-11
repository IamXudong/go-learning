// Package graphic2d 提供2d绘图相关实现
package graphic2d

// Point 点
type Point struct{ X, Y int }

// Circle 圆
type Circle struct {
	Center Point
	Radius int
}

// Wheel 车轮形
type Wheel struct {
	Circle Circle
	Spokes int
}

// PointScale 点缩放
func PointScale(p *Point, factor int) *Point {
	return &Point{p.X * factor, p.Y * factor}
}

// Circle1
type Circle1 struct {
	Point
	Radius int
}

// Wheel1
type Wheel1 struct {
	Circle1
	Spokes int
}

// point2
type point2 struct{X, Y int}

// Circle2
type Circle2 struct {
	point2
	Spokes int

}
