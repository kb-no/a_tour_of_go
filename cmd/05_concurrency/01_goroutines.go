// goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッド。

// go f(x, y, z)
// と書けば、新しいgoroutineが実行される。

// f(x, y, z)
// f , x , y , z の評価は、実行元(current)のgoroutineで実行され、 f の実行は、新しいgoroutineで実行される。

// goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある。
// sync パッケージは同期する際に役に立つ方法を提供しているが、別の方法があるためそれほど必要ない。

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}