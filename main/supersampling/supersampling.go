package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func main() {
	// 打开图片文件
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// 解码&读取图片象素
	pixels, err := getPixels(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// 采样生成新图片
	width, height := len(pixels[0]), len(pixels)
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xr := x + 1
			if xr >= width {
				xr = x
			}

			xl := x - 1
			if xl < 0 {
				xl = x
			}

			yt := y - 1
			if yt < 0 {
				yt = y
			}

			yb := y + 1
			if yb >= height {
				yb = y
			}
			p0 := pixels[y][x]
			p1 := pixels[yt][x]
			p2 := pixels[yt][xr]
			p3 := pixels[y][xr]
			p4 := pixels[yb][xr]
			p5 := pixels[yb][x]
			p6 := pixels[yb][xl]
			p7 := pixels[y][xl]
			p8 := pixels[yt][xl]

			newPixel := superSampling(p0, p1, p2, p3, p4, p5, p6, p7, p8)
			newImg.Set(x, y, color.RGBA(newPixel))
		}
	}

	png.Encode(os.Stdout, newImg)
}

func getPixels(f io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixel := Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
			row = append(row, pixel)
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

func superSampling(p0, p1, p2, p3, p4, p5, p6, p7, p8 Pixel) Pixel {

	r := int(p0.R)*4 + (int(p1.R)+int(p3.R)+int(p5.R)+int(p7.R))*2 + (int(p2.R)+int(p4.R)+int(p6.R)+int(p8.R))*1
	g := int(p0.G)*4 + (int(p1.G)+int(p3.G)+int(p5.G)+int(p7.G))*2 + (int(p2.G)+int(p4.G)+int(p6.G)+int(p8.G))*1
	b := int(p0.B)*4 + (int(p1.B)+int(p3.B)+int(p5.B)+int(p7.B))*2 + (int(p2.B)+int(p4.B)+int(p6.B)+int(p8.B))*1
	a := int(p0.A)*4 + (int(p1.A)+int(p3.A)+int(p5.A)+int(p7.A))*2 + (int(p2.A)+int(p4.A)+int(p6.A)+int(p8.A))*1

	r8 := uint8(r / 16)
	g8 := uint8(g / 16)
	b8 := uint8(b / 16)
	a8 := uint8(a / 16)

	return Pixel{r8, g8, b8, a8}
}
