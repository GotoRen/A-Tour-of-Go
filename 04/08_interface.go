/* interface */
package main

import (
"fmt"
"math"
)

/* 構造体 */
/*********************************/
type MyFloat float64

type Vertex struct {
	X, Y float64
}
/*********************************/

/* メソッド */
/*********************************/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
/*********************************/

/* インターフェース */
/*********************************/
type Abser interface {
	Abs() float64
}
/*********************************/

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	fmt.Println("f: ", f)
	a = f  // MyFloatはAbserを実装する
	fmt.Println("a: ", a)
	fmt.Println("a.Abs(): ", a.Abs())

	fmt.Println()

	v := Vertex{3, 4}
	fmt.Println("v: ", v)
	a = &v // VertexはAbserを実装する
	fmt.Println("a(&v): ", a)
	fmt.Println("a.Abs(): ", a.Abs())
}
/* 出力
> f:  -1.4142135623730951
> a(f):  -1.4142135623730951
> a.Abs():  1.4142135623730951

> v:  {3 4}
> a(&v):  &{3 4}
> a.Abs():  5
*/
