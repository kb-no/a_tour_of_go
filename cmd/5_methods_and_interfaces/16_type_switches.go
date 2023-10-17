// 型switch はいくつかの型アサーションを直列に使用できる。
// 型switchは通常のswitch文と似ているが、型switchのcaseは型(値ではない)を指定し、
// それらの値は指定されたインターフェースの値が保持する値の型と比較される。

// switch v := i.(type) {
// case T:
//     // here v has type T
// case S:
//     // here v has type S
// default:
//     // no match; here v has the same type as i
// }

// 型switchの宣言は、型アサーション i.(T) と同じ構文を持つが、特定の型 T はキーワード type に置き換えられる。
// このswitch文は、インターフェースの値 i が 型 T または S の値を保持するかどうかをテストする。
// T および S の各caseにおいて、変数 v はそれぞれ 型 T または S であり、i によって保持される値を保持する。
// defaultの場合(一致するものがない場合)、変数 v は同じインターフェース型で値は i となる。

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("The boolean %v\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}