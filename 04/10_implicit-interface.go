package main

import "fmt"

type INTERFACE interface {
	METHOD()
}

type TYPE struct {
	STRING string
}

// 暗黙的にタイプTYPEがインターフェースINTERFACEを実装する（明示的に宣言する必要はない）
func (t TYPE) METHOD() {
	fmt.Println(t.STRING)
}

func main() {
	var i INTERFACE // 変数INTERFACEはインターフェース型
	i = TYPE{"Hello, World!"}
	i.METHOD()
}
