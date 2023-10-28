// 初期化と後処理ステートメントの記述は任意。
// 必要なければ書かなくて良い。

package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 10; { // i := 0 や i++ などは今回for文内では不要なので書かなくて良い
		sum += sum
	}
	fmt.Println(sum)
}
