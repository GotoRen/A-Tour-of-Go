/* Exercise: Error */
package main

import (
	"fmt"
	"math"
)

/* 構造体 */
type ErrNegativeSqrt float64
/* Error */
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("エラー: %v", float64(e))
}
// ErrNegativeSqrt構造体はError()メソッドを実装しているため、fmt.Errorインターフェースとして扱えると見做す
/* 関数（引数: float64, 戻り値: float64, error） */
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // エラーの時、Error() 内を実行する
	}
	z := math.Sqrt(x)
	return z, nil                    // エラーでなければ nil を返す
}

func main() {
	fmt.Println(Sqrt(2))  // 1.4142135623730951 <nil>
	fmt.Println(Sqrt(-2)) // 0 エラー: -2
}
