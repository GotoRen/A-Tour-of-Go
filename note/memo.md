- 引数
  - 型が違う
  ```
  func add(x int, y float64) int {
	   return x + y
  }
  ```
  - 型が同じ
  ```
  // 省略可
  func add2(x, y int) int {
	   return x + y
  }
  ```
- 基本型（組み込み型）
  - `bool`
  - `string`
  - `int`
    - `int8  `
    - `int16  `
    - `int32  `
    - `int64`
  - `uint`
    - `uint8`
    - `uint16`
    - `uint32`
    - `uint64`
    - `uintptr`
  - `byte` // uint8 の別名
  - `rune` // int32 の別名（Unicode のコードポイントを表す）
  - `float`
    - `float32`
    - `float64`
  - `complex`
    - `complex64`
    - `complex128`
- 表示
  - `%T`：型
  - `%v`：値
  - `%d`：int, uint, ...
  - `%g`：float64, complex128, ...
- メモ
  - 定数は`:=`を使って宣言できない
  - 文の接続はコンマ`,`で行う
- <font color=Blue>`defer`</font>
  - `defer`へ渡した関数の実行を呼び出し元の関数の終わり（returnする）まで遅延させる
  - `defer`へ渡した関数の引数はすぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない
  - `defer`へ渡した関数が複数ある場合、その呼び出しはスタックされる
  → LIFO（Last In First Out）
- <font color=Red>ポインタ</font>
  - `*p`：変数pのポインタ
  - `nil`：ゼロ値
  → `var p *int`
  - `p = &i`：ポインタpへのオペレータ
- <font color=Red>構造体</font>
  - `struct`フィールドはドット`.`を用いてアクセスする
- <font color=Red>配列</font>
  - __固定長__
  - `var a [10]string`：string型の10個の変数の配列a
  - `b := [10]int{}`：int型の10個の変数の配列b
  - 初期値は0
- <font color=Red>スライス</font>
  - __可変長__
  - `a []int`：int型のスライスa
  - `a[low : high]`：インデックス`low`と`high`の境界を指定することによってスライスを形成
  - 配列への参照のようなもの
  - `[3]bool{true, true, false}` // 配列リテラル
  - `[]bool{true, true, false}` //上記と同様の配列を作成し、それを参照するスライス
