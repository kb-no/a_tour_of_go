// チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道。

// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる
// (データは、矢印の方向に流れる)

// マップとスライスのように、チャネルは使う前に以下のように生成:
// ch := make(chan int)
// 通常、片方が準備できるまで送受信はブロックされる。
// これにより、明確なロックや条件変数がなくても、goroutineの同期を可能にする。

// サンプルコードは、スライス内の数値を合算し、2つのgoroutine間で作業を分配する。
// 両方のgoroutineで計算が完了すると、最終結果が計算される。

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}