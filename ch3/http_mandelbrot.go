package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"math/cmplx"
	"os"
    "strconv"
)

func main() {
	http.HandleFunc("/display", display)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func display(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var m int = 2
    if len(r.Form["num"]) > 0 {
        num := r.Form["num"][0]
        if num != "" {
            m, _ = strconv.Atoi(num)
        }
    }
    if m > 10 {
        m = 10
    }
	const (
		xmin,ymin,xmax,ymax = -2, -2, +2, +2
		width, height       = 1024, 1024
	)
    imgfile, _ := os.Create("./http_mandelbrot.png")
    defer imgfile.Close()

	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py :=0; py<height; py++ {
		y := float64(py)/height*(ymax-ymin)+ymin
		for px :=0; px <width; px++ {
			x := float64(px)/width*(xmax-xmin)+xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(m ,z))
		}
	}
	png.Encode(imgfile, img)
	png.Encode(w, img)
}

func mandelbrot(m int, z complex128) color.Color {
	const iterations = 200
	const contrast   = 15
    nm := complex(float64(m), 0)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = cmplx.Pow(v, nm) + z
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
            } else  {
			    return color.Gray{255-contrast*n}
            }
        }
	}
	return color.Black
}
