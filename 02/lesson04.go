/* Exercise: loop and function */
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for d := 1.0; d*d > 1e-10; z -= d { // 1^-10
		d = (z*z - x) / (2 * z)
		fmt.Println(d)
		/**
		-0.5 = (1*1 - 2) / (2 * 1)
		0.08333333333333333 = (1.5*1.5 - 2) / (2 * 1.5)
		0.002450980392156932 = (1.417*1.417 - 2) / (2 * 1.417)
		...
		 */
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
