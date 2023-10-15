// nil インターフェースの値は、値も具体的な型も保持しない。
// 呼び出す 具体的な メソッドを示す型がインターフェースのタプル内に存在しないため、
// nil インターフェースのメソッドを呼び出すと、ランタイムエラーになる。

package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}