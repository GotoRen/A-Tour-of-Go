/* struct literal(定数) */
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // v1は構造体Vertexを持つ
	v2 = Vertex{X: 3}  // Yは暗黙的に0
	v3 = Vertex{}      // X, Yは暗黙的に0
	p  = &Vertex{1, 2} // pは構造体Vertexのポインタ
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
