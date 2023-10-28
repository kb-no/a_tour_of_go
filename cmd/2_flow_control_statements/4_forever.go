// ループ条件を省略すれば、無限ループ( infinite loop )になるで、無限ループをコンパクトに表現できる。

package main

import "fmt"

func main() {
	for {
		fmt.Println("loop")
	}
}

