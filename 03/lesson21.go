/* Exercise: Fibonacci closure */
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

//var a, b int = 1, 0
func fibonacci() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		//fmt.Println("a = ", a)
		//fmt.Println("b = ", b)
		return a
	}
}

func main() {
	f := fibonacci() // クロージャf f() = fibonacci
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
