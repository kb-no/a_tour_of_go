// structリテラルは、フィールドの値を列挙することで新しいstructの初期値の割り当てを示している。

// Name: 構文を使って、フィールドの一部だけを列挙することができる(この方法でのフィールドの指定順序は関係ない)。
// 例では X: 1 として X だけを初期化している。

// & を頭に付けると、新しく割り当てられたstructへのポインタを戻す。

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // なにも指定しないフィールドは、0が初期値となる。
	v3 = Vertex{}
	p  = &Vertex{1, 2} // Vertex型のポインタ
)
func main() {
	fmt.Println(v1, v2, v3, p)
}