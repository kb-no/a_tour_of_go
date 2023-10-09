// ここで、 Abs と Scale メソッドは関数として書きなおしてある。
// 再度、* を消してみる。 なぜ振る舞いが変わったのかわかる？ コンパイルするために、さらに何が必要だろうか。

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) { // *を消すとエラーになる
	fmt.Println(&v.X)
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(&v.X)
	Scale(&v, 10)
	fmt.Println(Abs(v))
}