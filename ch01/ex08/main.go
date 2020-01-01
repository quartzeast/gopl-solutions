package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// fetch 发送获取指定 url 的 html 页面并在终端打印
func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fetch(url)
		} else {
			fetch(url)
		}
	}
}
