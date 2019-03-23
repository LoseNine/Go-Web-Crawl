package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func regexp_data(html string) {
	file, err := os.Create("result.txt")
	checkerr(err)
	var data = make([]string, 2)
	name := regexp.MustCompile(`<em><font class="skcolor_ljg">(.*?)</em>`)
	data = name.FindAllString(html, -1)
	for _, i := range data {
		file.WriteString(i + "\n")
	}
}

func get_data(client http.Client, key string, n int, stops chan int) {
	url := "https://search.jd.com/Search?keyword=keywords&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=python&page=pages"
	url = strings.Replace(url, "keywords", key, -1)
	url = strings.Replace(url, "pages", strconv.Itoa(n), -1)
	res, err := http.NewRequest("GET", url, nil)
	res.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.84 Safari/537.36")
	checkerr(err)
	resp, err := client.Do(res)
	checkerr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	html := string(body)
	fmt.Println(html)
	regexp_data(html)
	stops <- 1
}

func crawl(key string, page int) {
	n := 1
	stops := make(chan int)
	client := http.Client{}
	for i := 0; i < page; i += 1 {
		go get_data(client, key, n, stops)
		n += 2
	}
	for i := 0; i < page; i++ {
		<-stops
	}
}

func main() {
	var key string
	var page int
	fmt.Println("输入关键词")
	fmt.Scan(&key)
	fmt.Println("输入页码")
	fmt.Scan(&page)
	crawl(key, page)
}
