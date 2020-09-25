/* pointer receiver */
package main

import (
"fmt"
"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバを持つSclaeメソッドを定義
// レシーバが指す変数vを変更できる
// mainで宣言したVertex変数を変更するためには、Scaleメソッドはポインタレシーバにする必要がある
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Printf("%f, %f\n", v.X, v.Y) // 30.000000, 40.000000
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

/***
> *Vertexのとき -> 50
>  Vertexのとき -> 5
***/
