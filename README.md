# Tour_go
- Golang practice
- _Command_
  - `$ go env`
  - `$ which go` 
- _Sentence_
  - 環境
  ```
  package main
  import (
      "fmt"
      "runtime"
  )
  func main() { 
      fmt.Print(runtime.GOOS + "\n") // 実行環境（OS）
  }
  ```
- _run_ 
  - go build コマンド
    - 実行ファイルを作成
      - `$ go build hoge.go`
    - 実行
      - `./hoge`
  - go run コマンド
    - コンパイルと実行を続けて実行
      - `$ go run hoge.go`
    - スクリプト言語のように開発できる
