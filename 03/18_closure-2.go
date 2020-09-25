/* closer その2 */
package main

import "fmt"

func adder() func(int) int {
	sum := 0 // 最初の一度だけ呼び出される
	return func(x int) int {
		sum = sum + x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
