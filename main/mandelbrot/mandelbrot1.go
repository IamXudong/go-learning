// mandelbrot函数生成一个png格式的mandelbrot分形图
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, xmax, ymin, ymax = -2, +2, -2, +2 // 坐标范围
		width, height          = 1024, 1024     // 图象尺寸
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// 扫描象素
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// 转换坐标
			x := (float64(px)/width)*(xmax-xmin) + xmin
			y := (float64(py)/height)*(ymax-ymin) + ymin
			z := complex(x, y)
			// 绘制象素
			img.Set(px, py, mandelbrot(z))
		}
	}
	// 编码输出
	png.Encode(os.Stdout, img)
}

// 计算象素点色值
func mandelbrot(z complex128) color.Color {
	const iterations = 200 // 迭代次数
	const contrast = 15    // 颜色偏差

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z // 复数z的平方与自身的和
		if cmplx.Abs(v) > 2 {
			r := (0xff - contrast*n)
			g := (0xff - contrast*n)
			b := contrast * n
			return color.RGBA{r, g, b, 0xff}
		}
	}
	// 默认返回蓝色
	return color.RGBA{0x00, 0x00, 0xff, 0xff}
}
