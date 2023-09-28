// var 宣言では、変数毎に初期化子( initializer )を与えることができる。

// 初期化子が与えられている場合、型を省略できる。その変数は初期化子が持つ型になる。

package main

import (
	"fmt"
)

var i int = 42
var j int = 100

func main() {
	var c, python, java = true, false, "no"
	fmt.Println(i, j, c, python, java)
}