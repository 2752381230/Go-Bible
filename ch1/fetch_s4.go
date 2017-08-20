package main

import (
	"fmt"
	"os"
	"net/http"
	"net/url" // e1.10
	"io"
	//"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 存储数据作为 cache 使用
	for _, url := range os.Args[1:] {
		fname, ok := check_url(url)
		if !ok {
			url = "http://" + url
		}
		go fetch(url, ch, fname)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string, fname string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// e1.10
	f, err := os.Create(fname + ".html")
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("open %s with error: %v\n", url, err)
		return
	}
	// 不知道对错，但是能写数据到文件中，暂时先这么写吧
	nbytes, err := io.Copy(f, resp.Body)
	f.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s with error: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs %7d %s", secs, nbytes, url)
}

func check_url(url_s string) (fname string, ok bool) {
	if url_s == "" {
		return
	}
	if strings.HasPrefix(url_s, "http://") || strings.HasPrefix(url_s, "https://") {
		u, _ := url.Parse(url_s)
		return u.Host, true
	} else {
		u, _ := url.Parse("http://" + url_s)
		return u.Host, false
	}
}
