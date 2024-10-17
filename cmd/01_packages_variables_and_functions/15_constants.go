// 定数( constant )は、 const キーワードを使って変数と同じように宣言する。
// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使える。
// なお、定数は := を使って宣言できない。

package main

import "fmt"

const Pi = 3.14 // constは:=で書くことができない。つまり:=の場合は必ず変数ということ。

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}