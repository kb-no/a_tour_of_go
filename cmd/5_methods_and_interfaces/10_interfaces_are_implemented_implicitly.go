// 型にメソッドを実装していくことによって、インタフェースを実装(満た)する。
// インタフェースを実装することを明示的に宣言する必要はない( "implements" キーワードは必要ない)。
// 暗黙のインターフェースは、インターフェースの定義をその実装から切り離す。
// インターフェースの実装は、事前の取り決めなしにパッケージに現れることがある。

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}