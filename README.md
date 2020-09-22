# Tour_go
- Golang practice
- _Command_
  - `$ go env`
  - `$ which go` 
  - `$ go run hoge.go`
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
