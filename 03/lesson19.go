/* function value */
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	fmt.Println("hoge=", fn(3, 4))
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // sqrt(169)

	fmt.Println(compute(hypot)) // sqrt(25)
	fmt.Println(compute(math.Pow)) // pow(9)
}
