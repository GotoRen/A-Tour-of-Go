/* pointer */
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i                 // pはiを指す
	fmt.Println("*p =", *p) // *pを読む
	*p = 21                 // ポインターを介してiを設定
	fmt.Println(" i =", i)  // iの新しい値を見る

	p = &j                  // pはjを指す
	fmt.Println("*p =", *p) // *pを読む
	*p = *p / 37            //
	fmt.Println(" j =", j)  // see the new value of j
}
