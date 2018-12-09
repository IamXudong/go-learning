// surface 根据一个三维曲面函数计算并生成svg
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 以像素表示的画布大小
	cells         = 160                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴的范围
	xyscale       = width / 2 / xyrange // x或y轴上每个单位长度的像素
	zscale        = height * 0.4        // z轴上每个单位长度的像素
	angle         = math.Pi / 6         // x, y轴的角度(=30度)
)

var cos30, sin30 = math.Cos(angle), math.Sin(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: cyan; fill: white; stroke-width: 0.1' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			z := math.Max(math.Max(az, bz), math.Max(cz, dz))
			r, g, b := c(z)
			fmt.Println(z)
			fmt.Println(r, g, b)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%d, %d, %d)'/>\n", ax, ay, bx, by, cx, cy, dx, dy, r, g, b, r, g, b)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// 求出网格单元(i, j)的顶点坐标(x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算出曲面的高度z
	z := f(x, y)

	// 将(x, y, z) 等角投射到二维SVG绘图平面上, 坐标是(sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到(0, 0)的距离
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}

func c(z float64) (int, int, int) {
	z += 0.22
	r := int(z * (255 / 1.22))
	g := int((1.22 - z) * (100 / 1.22))
	b := int((1.22 - z) * (255 / 1.22))
	return r, g, b
}
