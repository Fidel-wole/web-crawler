package main

import (
	"fmt"
	"net/http"
	"sync"
	"regexp"
	"io"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	visited := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var crawl func(string, int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 {
			return
		}

		mu.Lock()
		if visited[url] {
			mu.Unlock()
			return
		}
		visited[url] = true
		mu.Unlock()

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}

	wg.Add(1)
	go crawl(url, depth)

	wg.Wait()
}

type RealFetcher struct{}

func (f RealFetcher) Fetch(url string) (string, []string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}

	bodyString := string(body)
	urls := extractUrls(bodyString)

	return bodyString, urls, nil
}

func extractUrls(body string) []string {
	urlRegex := regexp.MustCompile(`https?://[^\s"']+`)
	return urlRegex.FindAllString(body, -1)
}

func main() {
	realFetcher := RealFetcher{}
	Crawl("https://golang.org/", 2, realFetcher)
}
