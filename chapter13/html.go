package main

import (
	"fmt"
	"io"
	"net/http"
)

func responseSize(channel chan Page, url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	channel <- Page{Url: url, Size: len(body)}
}

func reportPanic() {
	p := recover()
	if p != nil {
		fmt.Println(p)
	}
}

type Page struct {
	Url  string
	Size int
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://golang.org"}
	defer reportPanic()
	for _, url := range urls {
		go responseSize(pages, url)
	}
	for i := 0; i < len(urls); i++ {
		page := <- pages
		fmt.Printf("%s : %d\n",page.Url, page.Size)
	}
}
