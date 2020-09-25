/* mutating map */
package main

import "fmt"

func main() {
	m := make(map[string]int)

	/* 要素の挿入 */
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	/* 要素の更新 */
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	/* 要素の削除 */
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	/* 要素の確認 */
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

