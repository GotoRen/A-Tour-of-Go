/* method argument 2 */
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッド
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 関数
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v: ", v)
	fmt.Println("v.Abs():", v.Abs()) // (&v).Abs()として解釈される
	fmt.Println("v: ", v)
	fmt.Println("AbsFunc(v):", AbsFunc(v))
	/*******************************************/
	p := &Vertex{4, 3}
	fmt.Println("p: ", p)
	fmt.Println("p.Abs():", p.Abs()) // (*p).Abs()として解釈される
	fmt.Println("p: ", p)
	fmt.Println("AbsFunc(*p):", AbsFunc(*p))
}
