package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

const (
	xmax, ymax, xmin, ymin = 4, 4, -4, -4
	width, height          = 8192, 8192
)

func main() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
			os.Exit(1)
		}
		var x, y, scale, contrast int = 2, 2, 200, 15
		for k, v := range r.Form {
			var err error
			switch k {
			case "x":
				x, err = strconv.Atoi(v[0])
			case "y":
				y, err = strconv.Atoi(v[0])
			case "scale":
				scale, err = strconv.Atoi(v[0])
			case "contrast":
				contrast, err = strconv.Atoi(v[0])
			}
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}
		}
		draw(w, x, y, scale, contrast)
	}
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}

func draw(w io.Writer, x int, y int, scale int, contrast int) {
	xmax, ymax, xmin, ymin := float64(x), float64(y), float64(-x), float64(-y)
	width, height := (xmax-xmin)*float64(scale), (ymax-ymin)*float64(scale)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := float64(0); py < height; py++ {
		y := py/height*(ymax-ymin) + ymin
		for px := float64(0); px < width; px++ {
			x := px/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(int(px), int(py), newton(z, contrast))
		}
	}
	png.Encode(w, img)
}

func newton(z complex128, contrast int) color.Color {
	const iterations = 0xffffff
	var v complex128 = z
	for n := uint32(0); n < uint32(iterations/contrast); n++ {
		if cmplx.Abs(cmplx.Pow(v, 4)-1) < 0.00000001 {
			base := n * uint32(contrast)
			var r, g, b uint8
			switch {
			case base < 0xff:
				r = uint8(0xff)
				g = uint8(base)
				b = uint8(0x00)
			case base < 0xff00:
				r = uint8(0xff - base>>8)
				g = uint8(0xff)
				b = uint8(0x00)
			case base < 0xffff:
				r = uint8(0x00)
				g = uint8(0xff)
				b = uint8(base)
			case base < 0xff0000:
				r = uint8(0x00)
				g = uint8(0xff00ff - base)
				b = uint8(0xff)
			case base < 0xff00ff:
				r = uint8(base >> 16)
				g = uint8(0x00)
				b = uint8(0xff)
			case base < 0xffffff:
				r = uint8(0xff)
				g = uint8(0x00)
				b = uint8(0xffffff - base)
			default:
				r = uint8(0x00)
				g = uint8(0x00)
				b = uint8(0x00)
			}
			return color.RGBA{r, g, b, 0xff}
		}
		v = (3*cmplx.Pow(v, 4) + 1) / (4 * cmplx.Pow(v, 3))
	}
	return color.Black
}
