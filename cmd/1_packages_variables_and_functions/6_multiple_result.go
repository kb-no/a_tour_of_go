// 関数は複数の戻り値を返すことができる。

package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(swap("hello", "world"))
}

// 2つ目の戻り値にエラー型を指定して、処理がうまくいけば2つ目の戻り値にnilを渡して、
// エラーが発生すればそのまま2つ目の戻り値にエラー型の値を渡すような使われ方が多い。