# Tour_go
- 【A Tour of Go】
- _GOPATH_
```go
$GOPATH
 |------------ bin
 |              |----- fuga // ビルドされた実行可能ファイルが入る
 |
 |------------ pkg
 |              |----- darwin_amd64 // ビルドされたパッケージが入る『pkg/GOARCH/pkgname.a』
 |                       |----- hoge.a
 |               
 |------------ src
                |----- fuga // 実行可能なGoのコード『src/cmdname/*.go』
                |        |----- main.go
                |
                |----- hoge // ライブラリのGoのコード『src/pkgname/*.go』
                         |----- hoge.go
```

- _Command_
  - `$ go env`
  - `$ which go` 
- _`get`コマンド_
  - `$ go get 取得先`
    - `/Users/USER/go/src`配下に展開する
  - `github.com/tenntenn/greeting`の取得
    - `$ go get github.com/tenntenn/greeting`
  - `golang.org/x/tour/..`の取得
    - `$ go get golang.org/x/tour/gotour`
- _env_
  ```go
  package main
  
  import (
      "fmt"
      "runtime"
  )
  
  func main() { 
      fmt.Print(runtime.GOOS + "\n") // 実行環境（OS）
  }
  ```
- _Run_ 
  - go build コマンド
    - 実行ファイルを作成
      - `$ go build hoge.go`
    - 実行
      - `$ ./hoge`
  - go run コマンド
    - コンパイルと実行を続けて実行
      - `$ go run hoge.go`
    - スクリプト言語のように開発できる
- _Reference_
  - [Golang辞典](http://www.tohoho-web.com/ex/golang.html#goroutines)
  - [Golangの関数まとめてみた](https://qiita.com/pei0804/items/dd8acfba3dfe32530717)
