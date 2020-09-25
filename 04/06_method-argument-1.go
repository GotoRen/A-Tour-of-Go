/* method argument 1 */
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// メソッド + ポインタ
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X, v.Y) // 6 8 // 12 9
}
// 関数　+ ポインタ
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) 		 //(&v).Scaleとして解釈される
	ScaleFunc(&v, 10) //（引数: ポインタ）

	p := &Vertex{4, 3}
	p.Scale(3)		// (*p).Scaleとして解釈される
	ScaleFunc(p, 10) //（引数: 変数）

	fmt.Println(v, p)   // {60 80} &{120 90}
}
