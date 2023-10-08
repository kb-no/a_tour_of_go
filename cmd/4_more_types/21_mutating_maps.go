// map m の操作を見ていく。

// m へ要素(elem)の挿入や更新:
// m[key] = elem

// 要素の取得:
// elem = m[key]

// 要素の削除:
// delete(m, key)

// キーに対する要素が存在するかどうかは、2つの目の値で確認:
// elem, ok = m[key]
// もし、 m に key があれば、変数 ok は true となり、存在しなければ、 ok は false となる。

// なお、mapに key が存在しない場合、 elem はmapの要素の型のゼロ値になる。

// Note: もし elem や ok をまだ宣言していなければ、次のように := で短く宣言できる:
// elem, ok := m[key]

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // m["Answer"]はゼロ値（intの場合は0）になる
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // m["Answer"]がゼロ値なのか、0を入れているのか判断するためにokを用いる
	fmt.Println("The value:", v, "Present?", ok)
}