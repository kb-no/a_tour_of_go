// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略できる。

package main

import "fmt"

func add(x, y int) int { // ← (x int, y int)の省略
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}