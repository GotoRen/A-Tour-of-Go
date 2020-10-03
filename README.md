# Tour_go
- 【A Tour of Go】
- GOPATH
$GOPATH<br>
 |------------ bin<br>
 |                 |----- fuga // ビルドされた実行可能ファイルが入る<br>
 |<br>
 |------------ pkg<br>
 |                 |----- darwin_amd64 // ビルドされたパッケージが入る『`pkg/GOARCH/<br>pkgname.a`』
 |                             |----- hoge.a<br>
 |<br>                          
 |------------ src<br>
                   |----- fuga // 実行可能なGoのコード『`src/cmdname/*.go`』<br>
                   |         |----- main.go<br>
                   |<br>
                   |----- hoge // ライブラリのGoのコード『`src/pkgname/*.go`』<br>
                             |----- hoge.go<br>

- _Command_
  - `$ go env`
  - `$ which go` 
- _Golang_
  - `golang.org/x/tour/..`の入手
    - `$ go get golang.org/x/tour/gotour`
  - 環境
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
