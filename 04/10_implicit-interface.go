/* implicit interface */
package main

import "fmt"

/* 構造体 */
type TYPE struct {
	STRING string
}

/* メソッド */
// 暗黙的にタイプTYPEがインターフェースINTERFACEを実装する（明示的に宣言する必要はない）
func (t TYPE) METHOD() {
	fmt.Println(t.STRING)
}

/* インターフェース */
type INTERFACE interface {
	METHOD()
}

func main() {
	var i INTERFACE // 変数INTERFACEはインターフェース型
	i = TYPE{"Hello, World!"} // 代入
	i.METHOD() // 出力
}
