/* append */
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)
	// 元の配列sが変数群を追加する際に容量が小さい場合は、動的により大きいサイズの配列を割り当て直す

	// 追加はnilスライスで動作
	s = append(s, 0)
	printSlice(s)

	// スライスは必要に応じて大きくなる（容量）
	s = append(s, 1)
	printSlice(s)

	// 一度に複数追加できる
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
