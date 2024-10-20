// io パッケージは、データストリームを読むことを表現する io.Reader インタフェースを規定している。
// Goの標準ライブラリには、ファイル、ネットワーク接続、圧縮、暗号化などで、このインタフェースの 多くの実装 がある。

// io.Reader インタフェースは Read メソッドを持つ:
// func (T) Read(b []byte) (n int, err error)
// Read は、データを与えられたバイトスライスへ入れ、入れたバイトのサイズとエラーの値を返す。
// ストリームの終端は、 io.EOF のエラーで返す。

// 下記コードは、 strings.Reader を作成し、 8 byte毎に読み出している。

package main

import (
	"fmt"
	"strings"
	"io"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}