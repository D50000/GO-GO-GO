package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// SageMap is safe to use concurrently.
type SafeMap struct {
	mu sync.Mutex
	v  map[string]bool
}

// Check the url crawled or not.
func (s *SafeMap) CheckAndMark(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.v[url] {
		return false
	}
	s.v[url] = true
	return true
}

// Crawl the url.
func Crawl(url string, depth int, fetcher Fetcher, visited *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done() // Reduce counter when "WaitGroup finish".

	if depth <= 0 {
		return
	}

	// Break when url visited.
	if !visited.CheckAndMark(url) {
		return
	}

	// Fetching
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)

	// Crawl children url by goroutine.
	for _, u := range urls {
		wg.Add(1) // Add counter when "WaitGroup invoked".
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
}

func main() {
	var wg sync.WaitGroup // Make sure all goroutine finish.
	visited := SafeMap{v: make(map[string]bool)}

	wg.Add(1)                                                  // Add counter when "WaitGroup invoked".
	go Crawl("https://golang.org/", 4, fetcher, &visited, &wg) // Execute with goroutine.

	wg.Wait() // Wait for "wg goroutine" done.
}

// Mock fetching data.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// Mock data.
var fetcher = fakeFetcher{
	"https://golang.org/": {
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": {
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": {
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": {
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
