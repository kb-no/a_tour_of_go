// Pointers
// Goはポインタを扱う。 ポインタは値のメモリアドレスを指す。

// 変数 T のポインタは、 *T 型で、ゼロ値は nil。

// var p *int
// & オペレータは、そのオペランド( operand )へのポインタを引き出す。

// i := 42
// p = &i
// * オペレータは、ポインタの指す先の変数を示す。

// fmt.Println(*p) // ポインタpを通してiから値を読みだす
// *p = 21         // ポインタpを通してiへ値を代入する
// これは "dereferencing" または "indirecting" としてよく知られている。

// なお、C言語とは異なり、ポインタ演算はない。

package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         // p is a pointer to i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j         // p is a pointer to j
	*p = *p / 37   // same as (*p) / 37
	fmt.Println(j)
}