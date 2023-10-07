// スライスのゼロ値は nil。
// nil スライスは 0 の長さと容量を持っており、何の元となる配列も持っていない。

package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!")
	}
}