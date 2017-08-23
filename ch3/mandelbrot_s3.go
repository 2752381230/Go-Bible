package main

import (
	"math/cmplx"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin,ymin,xmax,ymax = -2, -2, +2, +2
		width, height       = 1024, 1024
	)
    imgfile, _ := os.Create("./see.png")
    defer imgfile.Close()

	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py :=0; py<height; py++ {
		y := float64(py)/height*(ymax-ymin)+ymin
		for px :=0; px <width; px++ {
			x := float64(px)/width*(xmax-xmin)+xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(imgfile, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast   = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v +z
		if cmplx.Abs(v) > 2 {
            // 基本上都维持在 0-20 之间
			if n > 200 {
                return color.RGBA{255, 0, 0, 255}
            } else if n > 100 {
                return color.RGBA{0, 0, 255, 255}
            } else if n > 50 {
                return color.RGBA{0, 0, 255, 255}
            } else if n > 30 {
                return color.RGBA{255, 255, 0, 255}
            } else if n > 20 {
                return color.RGBA{255, 255, 255, 255}
            } else if n > 7 {
                return color.RGBA{0, 255, 0, 255}
            } else if n > 6 {
                return color.RGBA{0, 0, 255, 255}
            } else if n > 5 {
                return color.RGBA{255, 255, 0, 255}
            } else if n > 4 {
                return color.RGBA{255, 0xd7, 0, 255}
            } else if n > 3 {
                return color.RGBA{255, 0x83, 0xfa, 255}
            } else if n > 2 {
                return color.RGBA{255, 0x82, 0x47, 255}
            } else if n > 1 {
                return color.RGBA{255, 0x7f, 0x24, 255}
            } else {
                return color.Gray{255-contrast*n}
            }
		}
	}
	return color.Black
}
