// Goのプログラムは、パッケージ( package )で構成される。
// プログラムは main パッケージから開始される。
// このプログラムでは "fmt" と "math/rand" パッケージをインポート( import )している。

// 規約で、パッケージ名はインポートパスの最後の要素と同じ名前になる。
// 例えば、インポートパスが "math/rand" のパッケージは、 package rand ステートメントで始まるファイル群で構成される。

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Number is ", rand.Intn(10))
}