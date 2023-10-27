// チャネルは、 バッファ ( buffer )として使える。
// バッファを持つチャネルを初期化するには、 make の２つ目の引数にバッファの長さを与える:
// ch := make(chan int, 100)

// バッファが詰まった時は、チャネルへの送信をブロックする。 バッファが空の時には、チャネルの受信をブロックする。

// バッファが詰まるようにサンプルコードを変更し、何が起きるのかを確認。

package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 2がバッファ。こいつをいじってみる。
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}