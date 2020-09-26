/* error */
package main

import (
	"fmt"
	"time"
)
/* 構造体 */
type MyError struct {
	When time.Time
	What string
}

/* Error */
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

/* 関数（引数: なし, 戻り値: error型） */
func run() error {
	// MyError構造体はError()メソッドを実装しているため、fmt.Errorインターフェースとして扱えると見做す
	return &MyError{time.Now(), "it didn't work"}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("失敗", err)
	} else {
		fmt.Println("成功")
	}

	/*
	if err := run(); err != nil {
		fmt.Println("失敗", err)
	} else {
		fmt.Println("成功")
	}
	*/
}
