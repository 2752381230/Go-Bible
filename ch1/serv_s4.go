package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}
var cycles  = 5
const (
	whiteIndex = 0
	blackIndex = 1
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Method = %s, URL=%s, URL.Proto=%s, URL.Path = %q\n", r.Method, r.URL, r.Proto, r.URL.Path)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAdde = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
func lissajous(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	lissajous_image(w, r)
	mu.Unlock()
}

func lissajous_image(out io.Writer, r *http.Request) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	// 变换不同的 cycles 图像还不一致，挺诡异的，暂且不深入研究了
	line_colors := map[string][]color.Color {
		"green" : []color.Color{color.White, color.RGBA{0x00,0x80,0x00,0xff}},
		"blue"  : []color.Color{color.White, color.RGBA{0x00,0x00,0xff,0xff}},
		"red"   : []color.Color{color.White, color.RGBA{0xff,0x00,0x00,0xff}},
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles_str := r.Form.Get("cycles")
	if cycles_str != "" {
		cycles, _ = strconv.Atoi(cycles_str)
	}
	line_str := r.Form.Get("color")
	if line_str != "" {
		if colors, ok := line_colors[line_str]; ok {
			palette = colors
		}
	}


	freq := rand.Float64()*3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i:= 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img  := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
