// ゼロ個のメソッドを指定されたインターフェース型は、 空のインターフェース と呼ばれる:
// interface{}

// 空のインターフェースは、任意の型の値を保持できる。 (全ての型は、少なくともゼロ個のメソッドを実装している。)
// 空のインターフェースは、未知の型の値を扱うコードで使用される。
// 例えば、 fmt.Print は interface{} 型の任意の数の引数を受け取る。

package main

import "fmt"

func main() {
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}