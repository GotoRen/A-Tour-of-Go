/* type conversion */
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f) // float -> uint
	fmt.Println(x, y, z)

	//var i int = 42
	//var f float64 = float64(i) // error: fが初期化させていない
	//var u uint = uint(f)
	//fmt.Println(i, f)
    // or
	//i := 42
	//f := 1
	//f := float64(i) // error: fが初期化させていない
	//fmt.Println(i, f, u)
	var a int = 43
	var b float64 = 21
	b = float64(a) // OK
	fmt.Println(a, b)
}
