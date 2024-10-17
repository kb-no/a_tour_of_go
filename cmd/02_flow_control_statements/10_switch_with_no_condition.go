// 条件のないswitchは、 switch true と書くことと同じ。

// このswitchの構造は、長くなりがちな "if-then-else" のつながりをシンプルに表現できる。

package main

import (
	"fmt"
	"time"
)

func main() {
	jst := time.FixedZone("JST", 9*60*60)
	t := time.Now().In(jst)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}