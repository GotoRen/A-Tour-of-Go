### 関数
- 引数
  - 型が違う
  ```go
  func add(x int, y float64) int {
	   return x + y
  }
  ```
  - 型が同じ
  ```go
  // 省略可
  func add2(x, y int) int {
	   return x + y
  }
  ```

### 型  
- 基本型（組み込み型）
  - `bool`
  - `string`
  - `int`
    - `int8`
    - `int16`
    - `int32`
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

### 文法
- 表示
  - `%T`：型
  - `%v`：値
  - `%d`：int, uint, ...
  - `%g`：float64, complex128, ...
- 宣言
  - 定数は`:=`を使って宣言できない
  - 文の接続はコンマ`,`で行う

### Defer
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
  - 定義
  ```go
  type Vertex struct{
	   a, b int
  }
  func main(){
      s := Vertex(1, 2)
      fmt.Println(s.a, s.b)
  }
  ```

### 配列
- __固定長__
- `var a [10]string`：string型の10個の変数の配列a
- `b := [10]int{}`：int型の10個の変数の配列b
- 初期値は0

### スライス
- __可変長__
- `a []int`：int型のスライスa
- `a[low : high]`：インデックス`low`と`high`の境界を指定することによってスライスを形成
- 配列への参照のようなもの
- `[3]bool{true, true, false}` // 配列リテラル
- `[]bool{true, true, false}` //上記と同様の配列を作成し、それを参照するスライス
- 動的配列の作成
  - `make`関数
  - `a := make([]int, 5)`
  - 第2引数：lenght, 第3引数：capacity
  → 容量を省略すると「長さ=容量」
  ```go
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
  b = b[:cap(b)] // len(b)=5, cap(b)=5
  b = b[1:]      // len(b)=4, cap(b)=4
  ```

### メソッド
- `strings.Join()`：各文字列の間に区切り文字を入れて文字列スライスをマージする
- `appends`：スライスに新しい要素を追加する
```go
var s []int
s = append(s, 2) // 第2引数：追加要素
```
  - 元の配列sが変数群を追加する際に容量が小さい場合は、動的により大きいサイズの配列を割り当て直す
  - スライスは必要に応じて大きくなる（容量）

## Map
- 基本
  - `>> map[キーの型]値の型`
  - 例：）
  ```go
  m := make(map[strig]int)
  m["Answer"] = 42
  ```
- __mapリテラル__
  - <u>`Key`が必要</u>
  ```go
  type Vertex struct {
  	Lat, Long float64
  }
  var m = map[string]Vertex{
  	"Bell Labs": Vertex{
  		40.68433,
  		-74.39967,
  	},
  	"Google": Vertex{
  		37.42202,
  		-122.08408,
  	},
  }
  var n = map[string]Vertex{
  	"Bell Labs": {40.68433, -74.39967},
  	"Google":    {37.42202, -122.08408},
  }
  func main() {
  	// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
  	fmt.Println(m)
  	fmt.Println(n)
  }
  ```  
  - __mutating map__
    - mへ要素（elem）の挿入や更新
      - `m[ley] = elem`
    - 要素の取得
      - `elem = m[key]`
    - 要素の削除
      - `delete(m, key)`
    - 要素の確認
      - `elem, ok = m[key]`
        - `ok`が`true`：`m`に`key`が存在する
        - `ok`が`false`：`m`に`key`が存在しない
    - ※ メモ
      - `elem`や`ok`をまだ宣言していない場合は`:=`で短く宣言する
      - `elme, ok := m[key]`      

### クロージャ（closure）
- それ自身の外部から変数を参照する関数値
- 参照された変数へアクセスして変えることができる
→ その関数は変数へバインドされている
