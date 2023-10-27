// structのフィールドは、structのポインタを通してアクセスすることもできる。

// フィールド X を持つstructのポインタ p がある場合、
// フィールド X にアクセスするには (*p).X のように書くことができる。
// しかし、この表記法は大変面倒なので、Goでは代わりに p.X と書くこともできる。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(p)
	p.X = 4 // p.Xでアクセスしてもいいし、v.Xでアクセスもできる
	fmt.Println(v)
}