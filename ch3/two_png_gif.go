package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/png"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/display", display)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func display(w http.ResponseWriter, q *http.Request) {
	f, err := os.Open("snowa.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	g, err := png.Decode(f)
	if err != nil {
		fmt.Println(err)
	}

	f2, err := os.Open("snowb.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	g2, err := png.Decode(f2)
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Create("test.gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	p1 := image.NewPaletted(image.Rect(0, 0, 800, 600), palette.Plan9)

	draw.Draw(p1, p1.Bounds(), g, image.ZP, draw.Src) //添加图片

	p2 := image.NewPaletted(image.Rect(0, 0, 800, 600), palette.Plan9)
	draw.Draw(p2, p2.Bounds(), g2, image.ZP, draw.Src) //添加图片

	g1 := &gif.GIF{
		Image:     []*image.Paletted{p1, p2},
		Delay:     []int{3, 3},
		LoopCount: 0,
	}

	gif.EncodeAll(w, g1)
	gif.EncodeAll(f1, g1)
}
