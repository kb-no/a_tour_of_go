package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64 // float64という型はこのパッケージで定義されていないのでMyFloatという形で再定義する

// Method: Goにはクラスという概念がないが、型にメソッドを定義することができる。
//         ファンクションは単純に入ってきた値を処理するもの。メソッドはstructのメンバ変数を使える。
func (v Vertex) Abs() float64 { // 呼び出し方はv.Abs()
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 { // 呼び出し方はAbs(v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) { // レシーバーでポインタ（ここでは *Vertex）を受け取ることもできる
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v, 10)
	fmt.Println(Abs(v))

	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}