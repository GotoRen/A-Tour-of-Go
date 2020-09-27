/* channel 2 */
package main

import (
	"fmt"
	"time"
)

func funcA(chA chan <- string) {
	time.Sleep(3 * time.Second)
	chA <- "Finished"		 // チャネルにメッセージを送信する
	close(chA)		 		 // 送り手チャネルをクローズする
}

func main() {
	chA := make(chan string) // チャネルを作成する
	go funcA(chA)			 // チャネルをゴルーチンに渡す
	msg := <- chA			 // チャネルからメッセージを受信する
	fmt.Println(msg)
}
