/* channel 1 */
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	fmt.Println(s)
	for _, v := range s {
		sum += v
	}
	c <- sum // チャネルに値を送信する
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int) 	// チャネルの生成

	//fmt.Println(s[:len(s)/2])

	/* ゴルーチンによる並列実行
	 * 2つのgoroutine間で作業を分配してスライス内の数値を合算する
	 */
	go sum(s[:len(s)/2], c) // チャネルをゴルーチンに渡す (1+2+3=6)
	go sum(s[len(s)/2:], c) // チャネルをゴルーチンに渡す (4+5+6=15)

	x, y := <-c, <-c 		// チャネルからメッセージを受信する
	fmt.Println(x, y, x+y)  // 15 6 21
}
