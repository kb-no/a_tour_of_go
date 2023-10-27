// スライスするときは、それらの既定値を代わりに使用することで上限または下限を省略することができる。

// 既定値は下限が 0 で上限はスライスの長さ。

// 以下の配列において
// var a [10]int

// これらのスライス式は等価:
// a[0:10]
// a[:10]
// a[0:]
// a[:]

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4] // [3, 5, 7]
	fmt.Println(s)

	s = s[:2] // [3, 5]
	fmt.Println(s)

	s = s[1:] // [5]
	fmt.Println(s)
}