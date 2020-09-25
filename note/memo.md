## 型
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

## 文法
- フォーマット
  - `fmt.Printf("%表示", 値)`
- 表示
  - `%T`：型
  - `%v`：値
  - `%d`：int, uint, ...
  - `%g`：float64, complex128, ... 
- 宣言
  - 型の宣言
    - `var 変数 型 = 値`
  - まとめて宣言
  ```go
  var (
    変数 型 = 値
    変数 型 = 値
    変数 型 = 値
    ...
  )
  ```  
- 定数型（const）
  - フォーマット
    - `const 定数 = 値`
  - 例：）  
  ```go
  const World = "world"
  fmt.Println("Hello", World)
  ```
  - まとめて宣言
  ```go
  const (
    定数 = 値
    定数 = 値
    定数 = 値
    ...
  )
  ```
- ゼロ値
```go
var i int     // 0
var f float64 // 0
var b bool    // false
var s string  // ""
```
- 自動型取得
  - `変数 := 値`    
  - ※ <font color=Red><u>定数`const`は`:=`を使って宣言できない</u></font>
  - ※ <font color=Red><u>関数内部でのみ使用できる</u></font>
- 文字列と変数の接続    
  - 文の接続はコンマ`,`で行う  
- グローバル変数
  - 関数の外側で宣言された変数
  - グローバルスコープ
  - `var 変数 型`
- プライベート変数
  - 関数の内側で宣言された変数
  - プライベートスコープ
  ```go
  func main() {
    var 変数 型
  }
  ```

## 関数
- フォーマット
```
func 関数名(引数1 型, 引数2 型, ...) 戻り値の型 {
  return 戻り値
}
```
- 型の省略
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
- 戻り値無し
  - 例：）
  ```go
  func split(sum int) (x, y int) {
	  x = sum * 4 / 9 
	  y = sum - x     
	  return
  }
  ```

## for文
- フォーマット
```go
sum := 0
for i:=下限; i<上限; i++ {
	処理
}
```
- 反復条件に式を代入
  - フォーマット
  ```
  for i:=初期値; 条件; 式{
    処理
  }
  ```
  - 例：）
  ```go
  for d := 1.0; d*d > 1e-10; z -= d {
  	d = (z*z - x) / (2 * z)
  }
  ```
- `range`
```go
変数1 := []型{値1, 値2, 値3, 値4, ...}
for i, 変数2 := range 変数1 {
		処理
}
```
- `range`インデックス無し
```go
変数1 := []型{値1, 値2, 値3, 値4, ...}
for _, 変数2 := range 変数1 {
		処理
}
```

## while文
  - フォーマット
  ```go
  i := 初期値
  for ; 条件式; {
  	処理
  }
  ```

## if文
- フォーマット
```go
if 条件式 {
  処理
}
```
- ショートステートメント
  - 例：）
  ```go
  if v := math.Pow(x, n); v < lim {
    return v;
  }
  ```
  - ifの動き
    - `v := math.Pow(x, n);`
      - 計算結果を`v`に代入
    - `v < lim`
      - `v`が`lim`よりも小さい時
    - ※ `v`はif文内でのみ使用できる      

## switch文
- フォーマット
```go
switch 変数 := 条件; 変数 {
  case 条件1:
  case 条件2:
  ...
  default:
} 
```

## defer
  - `defer`へ渡した関数の実行を呼び出し元の関数の終わり（returnする）まで遅延させる
  - `defer`へ渡した関数の引数はすぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない
  - `defer`へ渡した関数が複数ある場合、その呼び出しはスタックされる<br>
  → LIFO（Last In First Out）

## ポインタ
- `*p`：変数pのポインタ
- `nil`：ゼロ値<br>
→ `var p *int`
- `p = &i`：ポインタpへのオペレータ

## 構造体
- `struct`フィールドはドット`.`を用いてアクセスする
- フォーマット
```go
type 構造体名 struct{
   メンバ 型
   メンバ 型
   ...
}
func main(){
    変数 := 構造体名(値, 値, ...)
    fmt.Println(変数.メンバ, 変数.メンバ...)
}
```

## 配列
- <font color=Blue>__固定長__</font>
- フォーマット
  - `var 配列名 [個数]型`
- 例：）
  - `var array [10]string`：string型の10個の変数の配列`array`
  - `b := [10]int{}`：int型の10個の変数の配列`b`<br>
  → この場合、初期値はすべて`0`

## スライス
- <font color=Blue>__可変長__</font>
  - 配列への参照のようなもの
- フォーマット
```go
変数1 := [個数]int{値1, 値2, 値3, 値4, ...}
var 変数2 []int = 変数1[low:high]
```
- 概要
  - `a []int`：int型のスライス`a`
  - `a[low:high]`：インデックス`low`と`high`の境界を指定することによってスライスを形成
- 例：）
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4] // [3 5 7]
```
- 配列との違い
  - `[3]bool{true, true, false} // 配列` 
  - `[]bool{true, true, false} // スライス`  
- スライスに直接代入の例
  - `s := []int{2, 3, 5, 7, 11, 13}`
- 【よく使う】構造体スライスの例
```go
s := []struct {
	i int
	b bool
}{
	{2, true},
	{3, false},
	{5, true},
	{7, true},
	{11, false},
	{13, true},
}
```
- 「長さ」と「容量」
  - すべて同じ
  ```go
  s = s[0:最大]
  s = s[:最大]
  s = s[0:]
  s = s[:]
  ```
- <font color=Blue>__動的配列の作成__</font>
  - `make`関数
    - `変数 := make([]型, 長さ, 容量)`
      - 第2引数：長さ, 第3引数：容量
      - 容量を省略すると「長さ = 容量」となる
  - 例：）      
  ```go
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
  b = b[:cap(b)]         // len(b)=5, cap(b)=5
  b = b[1:]              // len(b)=4, cap(b)=4
  ```
- <font color=Blue>__要素の追加__</font>
  - `append`関数
    - `append(変数, 追加要素)`
  - 例：）
  ```go
  var s []int						
  printSlice(s) // len=0, cap=0	
  s = append(s, 0) 	   	
  printSlice(s) // len=1, cap=1
  s = append(s, 1) 	   	
  printSlice(s) // len=2, cap=2
  s = append(s, 2, 3, 4)  
  printSlice(s) // len=5, cap=6
  ```
  - 追加のルール
    - <u>元の配列`s`が変数群を追加する際に容量が小さい場合は、動的により大きいサイズの配列を割り当て直す</u>
    - 追加は`nil`スライスで動作
    - 一度に複数`append`できる    

## Map
- 『キー』と『値』を関連付ける
- フォーマット
  - `変数 := map[キーの型]値の型{"キー1": 値1, "キー2": 値2, ...}`
  - <u>固定長　配列を生成する場合</u>
  ```go
  変数 := map[キーの型]値の型{
    "キー1": 値1, 
    "キー2": 値2, 
    ...
  }
  ```
  - 例：）
    ```go
    m := map[string]string{
	  	"go": "golang",
	  	"rb": "ruby",
	  	"js": "javascript",
	  }
	  fmt.Println(m) // map[go:golang js:javascript rb:ruby]
    ```
  - <u>可変長配列を生成する場合</u>
  ```go
  変数 := make(map[キーの型]値の型)
  ```
  - 例：）
    ```go
    languages := make(map[string]string)
    fmt.Println(languages)            // map[]
    /* 型に則って任意のマップを作成していく */
    languages["go"] = "golang"
    languages["rb"] = "ruby"
    languages["js"] = "javascript"
    fmt.Println(languages)            // map[go:golang js:javascript rb:ruby]
    ```
- mapの変更
  - フォーマット
  ```go
  変数1 := make(map[キーの型]値の型)

  /* 要素の挿入・更新 */
  変数1[キー] = 値
  /* 要素の削除 */
  delete(変数1, キー)
  /* 要素の確認 */
  変数2, 変数3 := 変数1[キー] // キー（要素）が存在する場合true / キー（要素）が存在しない場合false
  ```
  
## クロージャ（closure）
- 参照された変数へアクセスして変えることができる<br>
→ その関数は変数へバインドされている
- main内の関数を変数に代入して外部関数に渡せる<br>
→ main内のクロージャに外部からアクセスできる
- フォーマット
```go
func 関数名 func匿名関数 戻り値の型 {
  処理
  return 匿名関数 戻り値の型 {
    処理
    return 戻り値
  }
}
```
- 例：）①
  ```go
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
  ```
- 例：）②
  ```go
  // laterは「引数に文字列を取り、戻り値に文字列を返す関数」を返す関数
  func later() func(string) string {
  	// storeはクロージャ内で使われている変数と結び付いた変数
  	// クロージャと結びついている限り破棄されない
  	var store string
  	store = "hoge" // 最初の一度だけ呼び出される
  	fmt.Println("store: " + store) 	  // debug

	  return func(next string) string { // クロージャ
	  	fmt.Println("next: " + next)  //debug
	  	s := store 	 // storeから一つ前に呼び出された文字列を取り出す（mainで変数fを使用している限りstoreの状態は保持される）
		  store = next // 今受け取っている文字列をstoreに格納
		  return s	 // 一つ前の文字列を返す
	  } 
  }

  func main() {
  	f := later()
  	fmt.Println("a: " + f("Golang"))
  	fmt.Println("b: " + f("is"))
  	fmt.Println("c: " + f("awesome!"))
  }
  /* 出力
  > store:
  > next: Golang
  > a:
  > next: is
  > b: Golang
  > next: awesome!
  > c:is
  */
  ```

*****
## <font color=Magenta>メソッド</font>
- `strings.Join()`
  - <u>各文字列の間に区切り文字を入れて文字列スライスをマージする</u>
  - フォーマット
  ```go
  変数1 := []型{値1, 値2, 値3, 値4, ...}
  変数2 := strings.Join(変数1, "区切り文字")
  ```
  - 例：）
  ```go
  board := [][]string{
  	[]string{"_", "_", "_"},
  	[]string{"_", "_", "_"},
  	[]string{"_", "_", "_"},
  }
  
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"
  
  merge := strings.Join(board[i], " ")
  ```
