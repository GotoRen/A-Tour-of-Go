/* append */
package main

import "fmt"

func main() {
	// 元の配列sが変数群を追加する際に容量が小さい場合は、動的により大きいサイズの配列を割り当て直す
	var s []int						// len=0, cap=0
	printSlice(s)

	// 追加はnilスライスで動作
	s = append(s, 0) 	   	// len=1, cap=1
	printSlice(s)

	// スライスは必要に応じて大きくなる（容量）
	s = append(s, 1) 	   	// len=2, cap=2
	printSlice(s)

	// 一度に複数追加できる
	s = append(s, 2, 3, 4)  // len=5, cap=6
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
