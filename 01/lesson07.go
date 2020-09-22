/* initialization */
package main

import "fmt"

var i, j int = 1, 2

func main() {
	k := 3 // 暗黙的な型宣言
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
