package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1 with error: %v\n", err)
			continue;
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1 url: %s with error: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
