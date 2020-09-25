/* closer その1 */
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		//fmt.Println(x*x + y*y)
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 値をhypotに渡す
	/**
	hypot := func(5, 12) // 無名関数
		return math.Sqrt(5*5 + 12*12)
		→ sqrt(169) = 13
	*/
	fmt.Println(compute(hypot))	    // クロージャhypotをcomputeに渡す
	/**
	fn = hypot
	return fn(3, 4)
	---------------
	hypot := func(3, 4) // 無名関数
	    return math.Sqrt(3*3 + 4*4)
		→ sqrt(25) = 5
	*/
	fmt.Println(compute(math.Pow))  // 関数computeに値を取りにいって累乗を計算する
	/**
	fn = math.Pow
	return fn(3, 4)
	math.Pow = 3^4 = 81
	 */
}
