/* buffered channel */
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // バッファ長を1に設定するとバッファが詰まって、以下の2つのチャネルでデットロックが発生する
	ch <- 1 // チャネル1
	ch <- 2 // チャネル2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
