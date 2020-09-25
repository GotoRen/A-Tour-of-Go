/* method */
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/* 匿名関数 */
func (v Vertex) Absolute() float64 { // 型にメソッドを定義できる
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Absolute())
}
