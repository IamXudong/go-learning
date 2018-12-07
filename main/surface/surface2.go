package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 120
	xyrange       = 40
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.05
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.5;' width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	m := int(x / (math.Pi * 2))
	n := int(y / (math.Pi * 2))

	dm := math.Abs(x - float64(m)*math.Pi*2)
	dn := math.Abs(y - float64(n)*math.Pi*2)

	if dm > math.Pi {
		dm = 2 * math.Pi - dm
	}

	if dn > math.Pi {
		dn = 2 * math.Pi - dn
	}

	r := math.Hypot(dm, dn)
	return math.Cos(r)
}
