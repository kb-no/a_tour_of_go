// interface(インタフェース)型は、メソッドのシグニチャの集まりで定義する。
// そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができる。

// 注意: この例は、※の部分でエラーになる。
// Abs メソッドが、 Vertex ではなく *Vertex の定義であり、
// Vertex が Abser インタフェースを実装していないということになるためエラーとなる。

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser ※ここでエラーになる

	// In the following line, v is a Vertex (not *Vertex)
	// and does not implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}