/* type switch */
package main

import "fmt"

func typeswitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%T, %v\n", v, v) // int, 21
	case string:
		fmt.Printf("%T, %q, %v byte\n", v, v, len(v)) // string, "Hello, World!", 13 byte
	default:
		fmt.Printf("%T\n, %v\n", v, v) // , true
	}
}

func main() {
	typeswitch(21)
	typeswitch("Hello, World!")
	typeswitch(true)
}
