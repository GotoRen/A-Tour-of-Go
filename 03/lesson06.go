/* slice to array */
package main

import "fmt"

func main() {
	names := [4]string{
		"hoge", // a[0]
		"fuga", // a[1]  b[0]
		"piyo", //       b[1]
		"tere",
	}
	fmt.Println(names)

	a := names[0:2] // スライス
	b := names[1:3] // スライス
	fmt.Println(a, b)

	b[0] = "XXX" // スライスb[0] = スライスa[1] = 配列names[1]
	fmt.Println(a, b)
	fmt.Println(names)
}
