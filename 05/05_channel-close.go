/* channel close */
package main

import (
	"fmt"
)

func calc(n int, c chan int) {
	x := 0
	for i := 0; i < n; i++ {
		c <- x
		x++
	}
	close(c) // これ以上、値を送信しないことを受信者に知らせるために送り手チャネルをクローズする
	         // クローズしない場合、受信者はnilチャネルから読み込もうとするためデットロックが発生する
}

func main() {
	ch := make(chan int, 10)
	go calc(cap(ch), ch)
	for i := range ch { // 送信チャネルが閉じられるまで、値を繰り返し受信し続ける
		fmt.Println(i)
	}
	/* チャネル状態の確認 */
	v, ok := <-ch
	fmt.Println("チャネル状態: ", v, ok) // 0 false
}
