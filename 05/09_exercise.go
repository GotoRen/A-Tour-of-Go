/* Exercise: Equivalent Binary Trees */
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)
/* treeパッケージの中身 */
/********
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
********/
// Walkはツリーのすべての値を送信する
// ツリーからチャネルへ
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	//fmt.Println("t.Value: ", t.Value)
	walk(t.Right, ch)
}

// 木が同じかどうかを決定する
// t1, t2には同じ値が含まれている
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		//fmt.Println("ok1: ", ok1)
		v2, ok2 := <-c2
		//fmt.Println("ok2: ", ok2)
		switch {
		case !ok1, !ok2:
			return ok1 == ok2
			//return true
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
