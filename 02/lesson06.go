/* defer */
package main

import "fmt"

func main() {
	//defer fmt.Println("World")
	//fmt.Println("Hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // stacking defer
	}
	fmt.Println("done")
}
