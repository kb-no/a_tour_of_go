// ポインタレシーバでメソッドを宣言できる。
// これはレシーバの型が、ある型 T への構文 *T があることを意味する。
// （なお、 T は *int のようなポインタ自身を取ることはできない）

// 例では *Vertex に Scale メソッドが定義されている。
// ポインタレシーバを持つメソッド(ここでは Scale )は、レシーバが指す変数を変更できる。
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的。

// Scale の宣言から * を消し、プログラムの振る舞いがどう変わるのかを確認してもOK。
// 変数レシーバでは、 Scale メソッドの操作は元の Vertex 変数のコピーを操作する。
// （これは関数の引数としての振るまいと同じ）。
// つまり main 関数で宣言した Vertex 変数を変更するためには、Scale メソッドはポインタレシーバにする必要がある。

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // *の有無による挙動の違いを確認
	fmt.Println(&v.X)
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(&v.X)
	v.Scale(10)
	fmt.Println(v.Abs())
}