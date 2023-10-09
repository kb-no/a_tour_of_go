// Goには、クラス( class )のしくみはないが、型にメソッド( method )を定義できる。
// メソッドは、特別なレシーバ( receiver )引数を関数に取る。
// レシーバは、 func キーワードとメソッド名の間に自身の引数リストで表現する。

// この例では、 Abs メソッドは v という名前の Vertex 型のレシーバを持つことを意味している。

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // Vertex型のインスタンス(仮にv)が実行できる関数
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}