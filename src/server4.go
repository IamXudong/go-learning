package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.Black, color.White}

const (
	blackIndex = 0
	whiteIndex = 1
)

var cycles int = 5
var res float64 = 0.001
var size int = 100
var nframes int = 64
var delay int = 8
var freq float64 = rand.Float64() * 3.0

func main() {
	handle := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		if len(r.Form["cycles"]) > 0 {
			cyclesf, err := strconv.Atoi(r.Form["cycles"][0])
			if err != nil {
				log.Print(err)
			}
			cycles = cyclesf
		}

		if len(r.Form["res"]) > 0 {
			resf, err := strconv.ParseFloat(r.Form["res"][0], 64)
			if err != nil {
				log.Print(err)
			}
			res = resf
		}

		if len(r.Form["size"]) > 0 {
			sizef, err := strconv.Atoi(r.Form["size"][0])
			if err != nil {
				log.Print(err)
			}
			size = sizef
		}

		if len(r.Form["nframes"]) > 0 {
			nframesf, err := strconv.Atoi(r.Form["nframes"][0])
			if err != nil {
				log.Print(err)
			}
			nframes = nframesf
		}

		if len(r.Form["delay"]) > 0 {
			delayf, err := strconv.Atoi(r.Form["delay"][0])
			if err != nil {
				log.Print(err)
			}
			delay = delayf
		}

		if len(r.Form["freq"]) > 0 {
			freqf, err := strconv.ParseFloat(r.Form["freq"][0], 64)
			if err != nil {
				log.Print(err)
			}
			freq = freqf
		}

		lissajous(w)
	}

	http.HandleFunc("/", handle)

	port := "8000"
	if len(os.Args[1:]) > 0 {
		port = os.Args[1]
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func lissajous(out io.Writer) {
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, size*2+1, size*2+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size-50)+0.5), size+int(y*float64(size-50)+0.5), whiteIndex)
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		phase += 0.1
	}

	gif.EncodeAll(out, &anim)
}
