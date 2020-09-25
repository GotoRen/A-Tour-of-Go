/* struct */
package main

import "fmt"

type Vertex struct {
	X int
	Y int
	// or: X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2} // 構造体を格納
	v.X = 4           // Xにアクセスして4を代入
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := &v // ポインタでアクセス
	p.X = 1e9
	fmt.Println(v)
}
