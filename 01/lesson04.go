/* multiple */
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	message := "Hello World!"
	fmt.Println(message)
	a, b := swap("Hello", "World!")
	fmt.Println(a, b)
}
