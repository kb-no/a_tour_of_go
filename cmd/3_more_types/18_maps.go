// map はキーと値とを関連付ける(マップする)。
// マップのゼロ値は nil。 nil マップはキーを持っておらず、またキーを追加することもできない。
// make 関数は指定された型のマップを初期化して、使用可能な状態で返す。

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}