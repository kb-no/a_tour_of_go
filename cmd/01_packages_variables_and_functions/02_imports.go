// このコードでは、括弧でパッケージのインポートをグループ化し、
// factoredインポートステートメント( factored import statement )としている。

package main

import (
	"fmt"
	"math"
)
// import "fmt"
// import "math"
// と書くこともできる。

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7)) // fmt.Printfはフォーマットで出力する
}