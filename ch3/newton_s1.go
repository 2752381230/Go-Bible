package main

import (
	"math/cmplx"
	"image"
	"image/color"
	"image/png"
	"os"
	//"fmt"
)

func main() {
	const (
		xmin,ymin,xmax,ymax = -2, -2, +2, +2
		width, height       = 1024, 1024
	)
    imgfile, _ := os.Create("./out_newton.png")
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
		v = cmplx.Pow(v, 4) - z
        //fmt.Println(cmplx.Abs(v))
        // 0-xxxx 数量相当大
		abs := cmplx.Abs(v)
		if abs > 2 {
            if abs > 50 {
                return color.RGBA{255, 0, 0, 255}
            } else if abs > 7 {
                return color.RGBA{255, 0, 255, 255}
            } else if abs > 6 {
                return color.RGBA{255, 255, 0, 255}
            } else if abs > 5 {
                return color.RGBA{0, 0, 255, 255}
            } else if abs > 4 {
                return color.RGBA{0, 255, 0, 255}
            } else if abs > 3 {
                return color.RGBA{255, 0, 0, 255}
            //} else if abs > 2 {
            //    return color.RGBA{255, 0, 0, 255}
            } else  {
			    return color.Gray{255-contrast*n}
            }
        }
    }
	return color.Black
}
