// if ステートメントは、 for のように、条件の前に、評価するための簡単なステートメントを書くことができる。

// ここで宣言された変数は、 if のスコープ内だけで有効。

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // if文の宣言と同時にvを定義する。vはifの中でのみ有効
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}