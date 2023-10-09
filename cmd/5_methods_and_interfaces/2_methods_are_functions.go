// メソッドは、レシーバ引数を伴う関数、だったよな？

// この Abs は、先ほどの例から機能を変えずに通常の関数として記述している。

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 { // Vertex型のインスタンス(仮にv)が実行できる関数
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}