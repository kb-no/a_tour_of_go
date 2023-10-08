// [n]T 型は、型 T の n 個の変数の配列( array )を表す。

// 以下は、intの10個の配列を宣言している:

// var a [10]int
// 配列の長さは、型の一部分。なので、配列のサイズを変えることはできない。

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}