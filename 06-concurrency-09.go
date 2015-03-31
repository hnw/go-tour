package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// ざっくり設計：
// ch(URL) -> 重複チェック(single) -> fetcher+parser(goroutine) -> ch(URL)

type CrawlUrl struct {
	Url   string
	Depth int
}

func concurrentFetcher(cu *CrawlUrl, ch chan *CrawlUrl, fetcher Fetcher) {
	body, urls, err := fetcher.Fetch(cu.Url)
	if err != nil {
		fmt.Println(err)
		ch <- nil
		return
	}
	fmt.Printf("found: %s %q\n", cu.Url, body)
	for _, u := range urls {
		ch <- &CrawlUrl{u, cu.Depth - 1}
	}
	ch <- nil
}

func cachedCrawler(ch chan *CrawlUrl, fetcher Fetcher) {
	alreadyFetched := make(map[string]bool)
	nFetcher := 0
	for cu := range ch {
		//fmt.Printf("cu=%v, nFetcher=%v\n", cu, nFetcher);
		if cu == nil {
			nFetcher--
			if nFetcher <= 0 {
				break
			}
		} else if cu.Depth > 0 && !alreadyFetched[cu.Url] {
			alreadyFetched[cu.Url] = true
			nFetcher++
			go concurrentFetcher(cu, ch, fetcher)
		}
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan *CrawlUrl, 10)
	ch <- &CrawlUrl{url, depth}
	cachedCrawler(ch, fetcher)
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
