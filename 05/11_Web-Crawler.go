/* Exercise: Web Crawler */
package main

import (
	"fmt"
	"sync"
)
/* インターフェース型: Fetcher
 * メンバ: Fetchメソッド
 */
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

/* struct型: fakeResult
 * メンバ: body(string), urls([]string)
 */
type fakeResult struct {
	body string
	urls []string
}

/* map型: fakeFetcher
 * キー: string, 値: struct
 */
type fakeFetcher map[string]*fakeResult

/* メソッド: Fetch
 * 自身のmapにURLをキーとする項目があるばそれを返し、無ければ"not found"を返す
 * レシーバ: fakeFetcher型 f
 * 引数: string型 url
 * 戻り値: string型, スライスstring型, error型
 */
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		//fmt.Println("debug-Fetch: ", res.body, res.urls)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

/* 関数: Crawl
 * fetcherを使用して再帰的にクロールする
 * クロールの開始点となるurl、そこから何階層辿るかの指定depth、そしてFetcher型を受け取る
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
	var wg sync.WaitGroup // WaitGroupの値を作る
	var crawl func(string, int) // クロージャ
	crawl = func(url string, depth int) {
		/* 階層depthが0以下ならその時点で終了
		 * そうでなければFetcherインターフェースに定義されているFetch()を引数のURLで実行
		 * 複数の戻り値をbody, urls, errで受ける
		 * エラーがnilでなければ、標準出力にエラーを記録して終了
		 */
		if depth <= 0 {
			return // ダミーデータ(fetcher)が何もなければその時点で終了
		}
		cache.Lock() // 排他制御開始
		if cache.visited[url] {   // visited map[url] にキーが存在しない場合
			cache.Unlock()
			return
		}
		cache.visited[url] = true // visited map[url] にキーが存在する場合
		cache.Unlock()


		/* Fetch()が問題なければ、取得したURLとbodyを出力してURLの詰まったsliceであるurlsをforとrangeで回す */
		/***************************************************/
		body, urls, err := fetcher.Fetch(url)
		//fmt.Println("debug", body, urls, err)
		if err != nil { // nilでない場合エラーメッセージを表示
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		//fmt.Println("err: ", err) // すべてnilを取得する
		//fmt.Println("debug-len(urls): ", len(urls))
		wg.Add(len(urls)) // 処理したい分wgをインクリメントする
		for _, u := range urls {
			//fmt.Println("debug-urls: ", urls)
			//fmt.Println("debug-u: ", u)
			//fmt.Println("debug-depth: ", depth)
			go func(u string) {
				crawl(u, depth-1) //depthをひとつ減らしてCrawl()を再帰実行
				wg.Done()
			}(u)
		}
		/***************************************************/
	}
	crawl(url, depth) // クロージャを実行
	wg.Wait() // すべてのゴルーチンの完了を待つ
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}


/* 変数: fetcher
 * 仮想フェッチャーURLの設定
 * ダミーデータを渡してfakeFetcherを取得
 * ダミーデータはURLをキーにして、bodyと次階層のurlsを持ったfakeResultを持つ
 */
var fetcher = fakeFetcher{
	// 1
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	// 2
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	// 3
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	// 4
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
