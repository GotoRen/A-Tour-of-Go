/* length and capacity */
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// すべて等価
	/*****************************/
	s = s[0:6]
	fmt.Println(s) // [5 7 11 13]
	s = s[:6]
	fmt.Println(s) // [5 7 11 13]
	s = s[0:]
	fmt.Println(s) // [5 7 11 13]
	s = s[:]
	/*****************************/

	s = s[1:4]     // 1~4
	fmt.Println(s) // [3 5 7]
	s = s[:2]      // 1~2
	fmt.Println(s) // [3 5]
	s = s[1:]      // 1~2
	fmt.Println(s) // [5]

	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2) // len=6 cap=6 [2 3 5 7 11 13]

	// スライスをスライス
	s2 = s2[:0]
	printSlice(s2) // len=0 cap=6 []
	// 長さを伸ばす
	s2 = s2[:4]
	printSlice(s2) // len=4 cap=6 [2 3 5 7]
	// 先頭2つを削除
	s2 = s2[2:] // len=2 cap=4 [5 7]
	printSlice(s2)
}

func printSlice(s2 []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)
}
