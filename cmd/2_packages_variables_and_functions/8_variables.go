// var ステートメントは変数( variable )を宣言する。
// 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できる。

// var ステートメントはパッケージ、または、関数で利用できる。

package main

import (
	"fmt"
)

var c, python, java bool // Goの場合は初期値を与えなくても初期値が型によって決められている（ゼロ値）。boolはfalse。

func main() {
	var i int // intは0が初期値。
	fmt.Println(i, c, python, java)
}