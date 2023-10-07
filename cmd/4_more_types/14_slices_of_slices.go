// スライスは、他のスライスを含む任意の型を含むことができる。

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a ticTacToe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X" // 左上
	board[2][2] = "O" // 右下
	board[1][2] = "X" // 中央右
	board[1][0] = "O" // 中央左
	board[0][2] = "X" // 右上

	// Print the board.
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}