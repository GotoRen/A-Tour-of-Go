/* Exercise: Web Crawler */
package main

import (
	"fmt"
	"sync"
)
/* Fetcherインターフェース
 * URLの本文を返す
 * そのページで見つかったURLのスライス
 */
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

/* Crawl関数
 * fetcherを使用して再帰的にクロールする
 * 開始位置 url から終了位置 depth まで
 * 引数: url(string), depth(int), fetcher(interface)
 * 戻り値: なし
 */
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := struct {
		visited map[string]bool
		sync.Mutex
	}{
		visited: make(map[string]bool),
	}
	var wg sync.WaitGroup
	var crawl func(string, int)
	crawl = func(url string, depth int) {
		if depth <= 0 {
			return
		}
		cache.Lock()
		if cache.visited[url] {
			cache.Unlock()
			return
		}
		cache.visited[url] = true
		cache.Unlock()
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		wg.Add(len(urls))
		for _, u := range urls {
			go func(u string) {
				crawl(u, depth-1)
				wg.Done()
			}(u)
		}
	}
	crawl(url, depth)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

/* struct型: fakeResult
 * メンバ: body(string), urls([]string)
 */
type fakeResult struct {
	body string
	urls []string
}

/* map型: akeFetcher
 * キー: string, 値: *fakeResult
 */
type fakeFetcher map[string]*fakeResult

/* メソッド: Fetch
 * レシーバ:
 */
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
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
