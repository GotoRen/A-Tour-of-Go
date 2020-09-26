/* empty interface */
package main

import "fmt"

func main() {
	var i interface{}   // 空のインターフェースを定義
	describe(i)

	i = 42      	    // 任意の値を追加
	describe(i)

	i = "Hello, World!" // 任意の値を追加
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
