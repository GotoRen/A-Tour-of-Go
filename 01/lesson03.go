/* function */
package main

import "fmt"

// 後ろに型名を書く
func add(x int, y int) int {
	return x + y
}

// 2追上の引数が同じ型の場合は省略可
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(12, 23))
}
