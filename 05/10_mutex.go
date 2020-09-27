/* mutex */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 構造体SafeCounterは同時に使用しても安全
type SafeCounter struct {
	mu sync.Mutex 	  // m: sync.Mutexライブラリ
	v  map[string]int // v: map型（キー: string型, 値: int型）
}

/* メソッド: Inc
 * 指定されたキーのカウンタをインクリメントする
 * レシーバ: struct型 c
 * 引数: string型 key
 * 戻り値: なし
 */
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// 一度に1つのゴルーチンだけがマップc.vにアクセスできるように排他制御
	c.v[key]++
	c.mu.Unlock()
}

/* メソッド: Value
 * 指定されたキーのカウンタの現在の値を返す
 * レシーバ: struct型 c
 * 引数: string型 key
 * 戻り値: int型 c.v[key]
 */
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// 一度に1つのゴルーチンだけがマップc.vにアクセスできるように排他制御
	defer c.mu.Unlock() // deferによって明示的にUnlockすることを示せる
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println("c.Value: ", c.Value("somekey"))
}
