// もっともよく使われているinterfaceの一つに fmt パッケージ に定義されている Stringer がある:
// type Stringer interface {
//     String() string
// }
// Stringer インタフェースは、stringとして表現することができる型。
// fmt パッケージ(と、多くのパッケージ)では、変数を文字列で出力するためにこのインタフェースがあることを確認する。

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 下記Stringerインタフェースをコメントアウトした時の出力と見比べる
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}