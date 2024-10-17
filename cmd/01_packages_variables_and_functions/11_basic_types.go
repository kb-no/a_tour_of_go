// Go言語の基本型(組み込み型)は次のとおり:

// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr : 符号なし整数。メモリサイズの指定やビット操作に用いられ、uintptrはポインタを扱う場面で使われる。
// byte（uint8 の別名）: uint8の別名で、主にバイナリデータや文字列の操作で使われる。
// rune （int32 の別名。Unicode のコードポイントを表す）: int32の別名で、Unicodeコードポイントを表す。マルチバイト文字（日本語など）の扱いで使用。
// float32 float64　: 浮動小数点数を扱う。計算精度が必要な場合はfloat64を使い、数値演算や科学計算で活用される。
// complex64 complex128 : 複素数を扱う。高度な数値解析や特殊なアルゴリズムで利用される。

// (訳注：runeとは古代文字を表す言葉(runes)だが、Goでは文字そのものを表すためにruneという言葉を使う。)
// 例では、変数宣言は、インポートステートメントと同様に、まとめて( factored )宣言可能。

// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bit。
// サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、
// 整数の変数が必要な場合は int を使うようにする。

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool       = false
	MaxInt uint64 = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}