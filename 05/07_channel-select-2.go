/* channel-select 2 */
package main

import "fmt"

func calc(c, quit chan int) {
	x := 0
	for {
		// 準備ができたcaseから実行される
		select {
		case c <- x: // 先にcが実行されて（チャネルを送信）
			x += 2
		case <-quit: // 次にquitが実行される（チャネルを受信）
			fmt.Println("セレクト実行は完了しました。\n")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)
	// ゴルーチンfunc() と calc が並列実行される
	go func() {
		// cから実行
		for i := 0; i < 10; i++ {
			fmt.Printf("i=%d: ch=%v\n", i, <-ch) // 関数からチャネルを受信して表示
		}
		// 次にqを実行
		quit <- 0 // 関数にチャネルを送信
	}()
	calc(ch, quit)
}
