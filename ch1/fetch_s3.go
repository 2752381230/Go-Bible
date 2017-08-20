package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 存储数据作为 cache 使用
	for _, url := range os.Args[1:] {
		if !check_url(url) {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s with error: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%0.2fs %7d %s", secs, nbytes, url)
}

func check_url(url_s string) (ok bool) {
	if url_s == "" {
		return
	}
	if strings.HasPrefix(url_s, "http://") || strings.HasPrefix(url_s, "https://") {
		return true
	}
	return false
}
