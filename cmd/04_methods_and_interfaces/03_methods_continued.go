// 例で挙げたstructの型だけではなく、任意の型(type)にもメソッドを宣言できる。

// 例は、 Abs メソッドを持つ、数値型の MyFloat 型。
// レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージにある必要がある。
// 他のパッケージに定義している型に対して、レシーバを伴うメソッドを宣言できない （組み込みの int などの型も同様）。

package main

import (
	"fmt"
	"math"
)

type MyFloat float64 // float64だとこのパッケージで使えないので一度myFloatに定義しなおしている。

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}