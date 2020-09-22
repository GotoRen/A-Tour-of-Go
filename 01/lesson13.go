/* numeric constant */
package main

import "fmt"

const (
	// 1ビットを100桁左にシフトして巨大な数を作成
	// 1の後に100個のゼロが続く2進数
	Big = 1 << 100
	// もう一度99桁右にシフトすると、1 << 1、または2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small) // 2
	fmt.Println(needInt(Small)) // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big)) // 1.2676506002282295e+29
}
