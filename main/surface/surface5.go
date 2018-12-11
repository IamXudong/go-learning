package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:"+os.Args[1], nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

const (
	width, height = 600, 320            // 画布长度与宽度
	cells         = 100                 // 网格数
	xyrange       = 30.0                // xy方向数值范围 -15 到 15
	xyscale       = width / 2 / xyrange // xy方向单位象素数量
	zscale        = height * 0.4        // z方向单位象素数量
	angle         = math.Pi / 6         // 投影角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: #eee; fill: white; stroke-width: 0.1' width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// 计算单元多边形顶点二维投射坐标
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// 根据最高顶点计算填充色
			z := math.Max(math.Max(az, bz), math.Max(cz, dz))
			color := c(z)

			// 输出多边形
			fmt.Fprintf(w, "<polygon fill='#%06x' points='%g,%g %g,%g %g,%g %g,%g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "%s\n", "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// 计算坐标
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算高度
	z := f(x, y)

	// 计算二维投影坐标
	ax := width/2 + (x-y)*cos30*xyscale
	ay := height/2 + (x+y)*sin30*xyscale - z*zscale

	return ax, ay, z
}

func f(x, y float64) float64 {
	// (x, y)点到坐标中心的距离
	r := math.Hypot(x, y)

	// 避免除数为0返回Inf
	if r == 0 {
		return 1
	}
	return math.Sin(r) / r
}

func c(z float64) uint32 {
	const cunit = 0xff / 1.22 // 单位色值
	z += 0.22                 // 高度转换为正值

	// 计算色值
	r := uint8(z * cunit)
	g := uint8(0)
	b := uint8((1.22 - z) * cunit)

	return uint32(r)<<16 + uint32(g)<<8 + uint32(b)
}
