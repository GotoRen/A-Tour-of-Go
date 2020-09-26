/* implicit interface */
package main

import (
	"fmt"
	"math"
)

/* 構造体 */
type T struct {
	S string
}
type F float64

/* メソッド */
// 暗黙的にタイプTがインターフェースIを実装する（明示的に宣言する必要はない）
// 変数の場合
/*
func (t T) M() {
	fmt.Println("t.S: ", t.S)
}
*/
// ポインタの場合
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("t.S: ", t.S)
}

func (f F) M() {
	fmt.Println("f: ", f)
}

/* インターフェース */
type I interface {
	M()
}

func main() {
	var i I // 変数iはインターフェース型
	/*******************************/
	var t *T
	i = t
	i.M()
	describe(i)
	fmt.Println()
	/*******************************/
	//i = T{"Hello, World!"} // 変数の場合
	i = &T{"Hello, World!"} // ポインタの場合
	fmt.Println("i: ", i)
	i.M()
	describe(i)

	fmt.Println()

	i = F(math.Pi)
	fmt.Println("i: ", i)
	i.M()
	describe(i)
}

/* タプル表示 */
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
