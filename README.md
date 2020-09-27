# Tour_go
- 【A Tour of Go】
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
- _Memo_
  - channel error
    - クローズしているチャネルに書き込むとパニックが起きる
    ```go
    func main() {
      ch := make(chan bool)
      close(ch)
      ch <- true
    }
    // panic: send on closed channel
    ```
    - クローズしているチャネルをクローズするとパニックが起きる
    ```go
    func main() {
      ch := make(chan bool)
      close(ch)
      close(ch)
    }
    // panic: close of closed channel
    ```
    - `nil`なチャネルに書き込むとデットロックが起きる
    ```go
    func main() {
      var ch chan bool
      ch <- true
    }
    // fatal error: all goroutines are asleep - deadlock!
    ```
    - `nil`なチャネルから読み込むとデットロックが起きる
    ```go
    func main() {
      var ch chan bool
      ch <- true
    }
    // fatal error: all goroutines are asleep - deadlock!
    ```
