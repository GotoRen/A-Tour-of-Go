/* pointer and function */
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	// ポインタを使用しない場合、関数内の処理はスコープの原理によってmain文に反映されない
	v.X = v.X * f // v.X = 30
	v.Y = v.Y * f // v.Y = 40
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // &v = {3, 4} // 関数Scaleで変数vのアドレスを書き換える（関数内の処理をmainに反映させる）
	fmt.Println(Abs(v))
}
