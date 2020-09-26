/* type assertion */
package main

import "fmt"

func main() {
	var i interface{} = "Hello, World"

	s := i.(string)
	fmt.Println(s)       // Hello, World

	s, ok := i.(string)
	fmt.Println(s, ok)   // Hello, World true
	/***************************************/
	//i = 12.345
	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)

	f = i.(float64) 	 // panic
	fmt.Println(f)
}
