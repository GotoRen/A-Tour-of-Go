# Tour_go
- 【A Tour of Go】
- _Command_
  - `$ go env`
  - `$ which go` 
- _Sentence_
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
      - `./hoge`
  - go run コマンド
    - コンパイルと実行を続けて実行
      - `$ go run hoge.go`
    - スクリプト言語のように開発できる
- _Reference_
  - [golang辞典](http://www.tohoho-web.com/ex/golang.html#goroutines)
  - [golangの関数まとめてみた](https://qiita.com/pei0804/items/dd8acfba3dfe32530717)
