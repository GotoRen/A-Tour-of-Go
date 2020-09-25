/* Exercise: closure */
package main

import "fmt"

// laterは「引数に文字列を取り、戻り値に文字列を返す関数」を返す関数
func later() func(string) string {
	// storeはクロージャ内で使われている変数と結び付いた変数
	// クロージャと結びついている限り破棄されない
	var store string
	store = "hoge" // 最初の一度だけ呼び出される
	fmt.Println("store: " + store) 	  // debug

	return func(next string) string { // クロージャ
		fmt.Println("next: " + next)  //debug
		s := store 	 // storeから一つ前に呼び出された文字列を取り出す（mainで変数fを使用している限りstoreの状態は保持される）
		store = next // 今受け取っている文字列をstoreに格納
		return s	 // 一つ前の文字列を返す
	}
}

func main() {
	f := later()
	fmt.Println("a: " + f("Golang"))
	fmt.Println("b: " + f("is"))
	fmt.Println("c: " + f("awesome!"))
}

/* 出力
> store:
> next: Golang
> a:
> next: is
> b: Golang
> next: awesome!
> c:is
*/
