// 関数は、0個以上の引数を取ることができる。
// この例では、 add 関数は、 int 型の２つのパラメータを取る。
// 変数名の 後ろ に型名を書くことに注意。

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}