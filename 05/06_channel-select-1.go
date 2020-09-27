/* channel select 1 */
package main

import (
	"fmt"
	"time"
)

/* 以下の関数はAが先に実行される */
func funcA(chA chan <- string) {
	time.Sleep(1 * time.Second)
	chA <- "funcA Finished"
}
func funcB(chB chan <- string) {
	time.Sleep(2 * time.Second)
	chB <- "funcB Finished"
}

func main() {
	// チャネル生成
	chA := make(chan string)
	chB := make(chan string)
	// チャネルクローズ
	defer close(chA)
	defer close(chB)
	// 終了フラグ
	finflagA := false
	finflagB := false
	// ゴルーチン（並列実行）
	go funcA(chA)
	go funcB(chB)

	for {
		/* ゴルーチン funcA() と funcB() 双方の待ち合わせ */
		// 準備ができたcaseから実行される
		select {
		case msg := <- chA: // 最初にAが実行されて
			finflagA = true
			fmt.Println(msg)
		case msg := <- chB: // 次にBが実行される
			finflagB = true
			fmt.Println(msg)
		}
		// funcA() と funcB()の実行が終了したら無限ループを抜ける
		if finflagA && finflagB {
			break
		}
	}
}
