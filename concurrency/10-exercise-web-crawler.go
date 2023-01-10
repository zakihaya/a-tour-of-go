package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
	IsFetched(url string) bool
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	time.Sleep(time.Second)
	if depth <= 0 {
		return
	}
	// Don't fetch the same url twice.
	if fetcher.IsFetched(url) {
		fmt.Printf("already crawled %s\n", url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found %s %q\n", url, body)
	wg.Add(len(urls))
	for _, u := range urls {
		// Fetch urls in parallel.
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher)
		}(u)

	}
	return
}

func main() {
	fmt.Println("main")
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
	fmt.Println("main end")
}

// FetcherFaker is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func (f fakeFetcher) IsFetched(url string) bool {
	mu.Lock()
	defer mu.Unlock()
	// _: value
	// ok: whether key exists
	_, ok := fetchedUrls[url]
	if ok {
		return true
	}
	fetchedUrls[url] = true
	return false
}

var (
	wg          sync.WaitGroup
	mu          sync.Mutex
	fetchedUrls = make(map[string]bool)
)
