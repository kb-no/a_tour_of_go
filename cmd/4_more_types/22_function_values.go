// 関数も変数。他の変数のように関数を渡すことができる。
// 関数値( function value )は、関数の引数に取ることもできるし、戻り値としても利用できる。

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 { // 関数を引数にとる関数
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}