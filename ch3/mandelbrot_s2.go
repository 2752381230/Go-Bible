package main

import (
    "math"
	"math/cmplx"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	const (
		xmin,ymin,xmax,ymax = -2, -2, +2, +2
		width, height       = 1024*4, 1024*4
	)
	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py :=0; py<height; py++ {
		y := float64(py)/height*(ymax-ymin)+ymin
		for px :=0; px <width; px++ {
			x := float64(px)/width*(xmax-xmin)+xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//const contrast   = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v +z
		if cmplx.Abs(v) > 2 {
            x := int(math.Trunc(math.Abs(real(v))))
            y := int(math.Trunc(math.Abs(imag(v))))
			return myfunc(x, y)
		}
	}
	return color.Black
}
func myfunc(a, b int) color.Color  {
    tmp := a+b*4096
    z   := tmp%256
    tmp = (tmp-z)/256
    y  := tmp%256
    x  := (tmp-y)/256
    return color.RGBA{uint8(x), uint8(y), uint8(z), 255}
}
