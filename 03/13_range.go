/* range */
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	fmt.Println("\n")

	// インデックス値はアンダーバーで代入を捨てることができる
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2^i
		fmt.Printf("%d: pow2 = %d\n", i, pow2[i])
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
