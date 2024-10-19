// もし、mapに渡すトップレベルの型が単純な型名である場合は、リテラルの要素から推定できるので、
// その型名を省略することができる。

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{ // ←ここにVertexと記載があるので、
	"Bell Labs": { // ←ここのVertexは省略可能。
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.083,
	},
}

func main() {
	fmt.Println(m)
}