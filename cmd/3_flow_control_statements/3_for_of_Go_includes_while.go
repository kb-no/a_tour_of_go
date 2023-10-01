// セミコロン(;)を省略することも可能。つまり、C言語などにある while は、Goでは for だけを使う。

package main

import (
	"fmt"
)

func main() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}