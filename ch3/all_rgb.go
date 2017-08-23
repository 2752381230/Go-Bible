package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
    //"fmt"
)

func main() {
	const (
		width, height = 1024*4, 1024*4
	)
    imgfile, _ := os.Create("./see.png")
    defer imgfile.Close()

	img := image.NewRGBA(image.Rect(0,0, width, height))
	for py :=0; py<height; py++ {
		for px :=0; px <width; px++ {
			img.Set(px, py, myfunc(px, py))
		}
	}
	png.Encode(imgfile, img)
}

func myfunc(a, b int) color.Color  {
    tmp := a+b*4096
    z   := tmp%256
    tmp = (tmp-z)/256
    y  := tmp%256
    x  := (tmp-y)/256
    //fmt.Println(a,b, x, y, z)
    return color.RGBA{uint8(x), uint8(y), uint8(z), 255}
}
