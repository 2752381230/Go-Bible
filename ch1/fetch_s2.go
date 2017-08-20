package main

import (
	"fmt"
	"os"
	"net/http"
	"io"      // e1.7
	"strings" // e1.8
)

func main() {
	for _, url := range os.Args[1:] {
		// e1.8
		if !check_url(url) {
			url = "http://"+url
		}
		//fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1 with error: %v\n", err)
			continue;
		}
		//e1.9
		fmt.Println(resp.Status)
		count, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1 url: %s with error: %v\n", url, err)
			os.Exit(1)
		}
		// 因为直接到 stdout了不需要在打印了
		fmt.Println(count)
	}
}

// e1.8
func check_url(url string) (ok bool) {
	if url == "" {
		return
	}
	if strings.HasPrefix(url, "http://") {
		return true
	} else if strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}
