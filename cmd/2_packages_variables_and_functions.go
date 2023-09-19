package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Random number is", rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(15))
	var i int // 初期値は0
	fmt.Println(i, c, python, java)

	// 変数 v 、型 T があった場合、 T(v) は、変数 v を T 型へ変換
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使用可能
	const world = "world" // constは:=で置き換えられない
	fmt.Println("Hello", world)
}

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool // 初期値はfalse