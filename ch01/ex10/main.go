// ex10 对每个 URL 执行两次请求，查看两次时间是否有较大的差别， 并且每次获取到的响应内容是否一致。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	results := make(map[string]string)
	for _, url := range os.Args[1:] {
		for i := 0; i < 2; i++ {
			go fetch(url, ch)
		}
	}
	for range os.Args[1:] {
		for i := 0; i < 2; i++ {
			s := <-ch
			url := s[strings.LastIndex(s, " ")+1:]
			results[s] = url
		}
	}

	// 对相同的 URL 的两次请求放在一起输出，方便比较
	for r, u := range results {
		u1 := u
		s := r
		fmt.Println(s)     // 输出遍历到的第一个结果
		delete(results, r) // 删除 map 中的该元素
		len := len(results)
		i := 0                        // 设置一个计数量 i，用来判断 if u1 == u2 {...} 是否执行过。
		for r1, u2 := range results { // 遍历删除上述元素后的 map，查找是否有相同的 URL。
			if u1 == u2 {
				fmt.Println(r1)     // 如果查找到相同的 URL，两次请求的结果不一样
				delete(results, r1) // 输出第二个结果，并从 map 中删除
				break               // break 掉循环
			}
			i++ //
		}
		if i == len { // 如果 i == len 说明， if u1 == u2 {...} 没有执行过，即没找到相同的 url，两次请求的结果相同，则要将 s 再输出一遍。
			fmt.Println(s)
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
