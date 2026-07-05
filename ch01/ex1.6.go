package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	lissajous16(os.Stdout)
}

func lissajous16(out io.Writer) {
	var palette = []color.Color{color.White, color.Black,
		color.RGBA{0x0, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0x0, 0x00, 0xff},
		color.RGBA{0x0, 0x0, 0xff, 0xff},
	}

	const (
		whiteIndex = 0
		blackIndex = 1
		greenIndex = 2
		redIndex   = 3
		blueIndex  = 4
	)

	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	phase := 0.0
	anim := gif.GIF{LoopCount: nframes}
	for colorIndex := greenIndex; colorIndex <= blueIndex; colorIndex++ {
		for i := 0; i < nframes; i++ {
			rect := image.Rect(0, 0, 2*size+1, 2*size+1)
			img := image.NewPaletted(rect, palette)
			// ex1.5. background color is black
			for x := 0; x < 2*size+1; x++ {
				for y := 0; y < 2*size+1; y++ {
					img.SetColorIndex(x, y, blackIndex)
				}
			}
			for t := 0.0; t < cycles*2*math.Pi; t += res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
	}
	gif.EncodeAll(out, &anim)
}
