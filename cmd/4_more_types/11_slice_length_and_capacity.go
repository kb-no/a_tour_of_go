// スライスは長さ( length )と容量( capacity )の両方を持っている。
// スライスの長さは、それに含まれる要素の数。
// スライスの容量は、スライスの最初の要素から数えて、元となる配列の要素数。

// スライス s の長さと容量は len(s) と cap(s) という式を使用して得ることができる。
// 十分な容量を持って提供されているスライスを再スライスすることによって、スライスの長さを伸ばすことができる。
// その容量を超えて伸ばしたときに何が起こるかを見るため、プログラム例でのスライスのいずれかの操作を変更してみる。

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]

	s = s[1:]
	printSlice(s) // len=1 cap=3 [7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}